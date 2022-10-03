package cosmosdbaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/cosmosdbaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account azurerm_cosmosdb_account}.
type CosmosdbAccount interface {
	cdktf.TerraformResource
	AccessKeyMetadataWritesEnabled() interface{}
	SetAccessKeyMetadataWritesEnabled(val interface{})
	AccessKeyMetadataWritesEnabledInput() interface{}
	AnalyticalStorage() CosmosdbAccountAnalyticalStorageOutputReference
	AnalyticalStorageEnabled() interface{}
	SetAnalyticalStorageEnabled(val interface{})
	AnalyticalStorageEnabledInput() interface{}
	AnalyticalStorageInput() *CosmosdbAccountAnalyticalStorage
	Backup() CosmosdbAccountBackupOutputReference
	BackupInput() *CosmosdbAccountBackup
	Capabilities() CosmosdbAccountCapabilitiesList
	CapabilitiesInput() interface{}
	Capacity() CosmosdbAccountCapacityOutputReference
	CapacityInput() *CosmosdbAccountCapacity
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionStrings() *[]*string
	ConsistencyPolicy() CosmosdbAccountConsistencyPolicyOutputReference
	ConsistencyPolicyInput() *CosmosdbAccountConsistencyPolicy
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CorsRule() CosmosdbAccountCorsRuleOutputReference
	CorsRuleInput() *CosmosdbAccountCorsRule
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreateMode() *string
	SetCreateMode(val *string)
	CreateModeInput() *string
	DefaultIdentityType() *string
	SetDefaultIdentityType(val *string)
	DefaultIdentityTypeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableAutomaticFailover() interface{}
	SetEnableAutomaticFailover(val interface{})
	EnableAutomaticFailoverInput() interface{}
	EnableFreeTier() interface{}
	SetEnableFreeTier(val interface{})
	EnableFreeTierInput() interface{}
	EnableMultipleWriteLocations() interface{}
	SetEnableMultipleWriteLocations(val interface{})
	EnableMultipleWriteLocationsInput() interface{}
	Endpoint() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeoLocation() CosmosdbAccountGeoLocationList
	GeoLocationInput() interface{}
	Id() *string
	SetId(val *string)
	Identity() CosmosdbAccountIdentityOutputReference
	IdentityInput() *CosmosdbAccountIdentity
	IdInput() *string
	IpRangeFilter() *string
	SetIpRangeFilter(val *string)
	IpRangeFilterInput() *string
	IsVirtualNetworkFilterEnabled() interface{}
	SetIsVirtualNetworkFilterEnabled(val interface{})
	IsVirtualNetworkFilterEnabledInput() interface{}
	KeyVaultKeyId() *string
	SetKeyVaultKeyId(val *string)
	KeyVaultKeyIdInput() *string
	Kind() *string
	SetKind(val *string)
	KindInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalAuthenticationDisabled() interface{}
	SetLocalAuthenticationDisabled(val interface{})
	LocalAuthenticationDisabledInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MongoServerVersion() *string
	SetMongoServerVersion(val *string)
	MongoServerVersionInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkAclBypassForAzureServices() interface{}
	SetNetworkAclBypassForAzureServices(val interface{})
	NetworkAclBypassForAzureServicesInput() interface{}
	NetworkAclBypassIds() *[]*string
	SetNetworkAclBypassIds(val *[]*string)
	NetworkAclBypassIdsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	OfferType() *string
	SetOfferType(val *string)
	OfferTypeInput() *string
	PrimaryKey() *string
	PrimaryReadonlyKey() *string
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
	// Experimental.
	RawOverrides() interface{}
	ReadEndpoints() *[]*string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Restore() CosmosdbAccountRestoreOutputReference
	RestoreInput() *CosmosdbAccountRestore
	SecondaryKey() *string
	SecondaryReadonlyKey() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CosmosdbAccountTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualNetworkRule() CosmosdbAccountVirtualNetworkRuleList
	VirtualNetworkRuleInput() interface{}
	WriteEndpoints() *[]*string
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
	PutAnalyticalStorage(value *CosmosdbAccountAnalyticalStorage)
	PutBackup(value *CosmosdbAccountBackup)
	PutCapabilities(value interface{})
	PutCapacity(value *CosmosdbAccountCapacity)
	PutConsistencyPolicy(value *CosmosdbAccountConsistencyPolicy)
	PutCorsRule(value *CosmosdbAccountCorsRule)
	PutGeoLocation(value interface{})
	PutIdentity(value *CosmosdbAccountIdentity)
	PutRestore(value *CosmosdbAccountRestore)
	PutTimeouts(value *CosmosdbAccountTimeouts)
	PutVirtualNetworkRule(value interface{})
	ResetAccessKeyMetadataWritesEnabled()
	ResetAnalyticalStorage()
	ResetAnalyticalStorageEnabled()
	ResetBackup()
	ResetCapabilities()
	ResetCapacity()
	ResetCorsRule()
	ResetCreateMode()
	ResetDefaultIdentityType()
	ResetEnableAutomaticFailover()
	ResetEnableFreeTier()
	ResetEnableMultipleWriteLocations()
	ResetId()
	ResetIdentity()
	ResetIpRangeFilter()
	ResetIsVirtualNetworkFilterEnabled()
	ResetKeyVaultKeyId()
	ResetKind()
	ResetLocalAuthenticationDisabled()
	ResetMongoServerVersion()
	ResetNetworkAclBypassForAzureServices()
	ResetNetworkAclBypassIds()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetRestore()
	ResetTags()
	ResetTimeouts()
	ResetVirtualNetworkRule()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CosmosdbAccount
type jsiiProxy_CosmosdbAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CosmosdbAccount) AccessKeyMetadataWritesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessKeyMetadataWritesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) AccessKeyMetadataWritesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessKeyMetadataWritesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) AnalyticalStorage() CosmosdbAccountAnalyticalStorageOutputReference {
	var returns CosmosdbAccountAnalyticalStorageOutputReference
	_jsii_.Get(
		j,
		"analyticalStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) AnalyticalStorageEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analyticalStorageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) AnalyticalStorageEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analyticalStorageEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) AnalyticalStorageInput() *CosmosdbAccountAnalyticalStorage {
	var returns *CosmosdbAccountAnalyticalStorage
	_jsii_.Get(
		j,
		"analyticalStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Backup() CosmosdbAccountBackupOutputReference {
	var returns CosmosdbAccountBackupOutputReference
	_jsii_.Get(
		j,
		"backup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) BackupInput() *CosmosdbAccountBackup {
	var returns *CosmosdbAccountBackup
	_jsii_.Get(
		j,
		"backupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Capabilities() CosmosdbAccountCapabilitiesList {
	var returns CosmosdbAccountCapabilitiesList
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) CapabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Capacity() CosmosdbAccountCapacityOutputReference {
	var returns CosmosdbAccountCapacityOutputReference
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) CapacityInput() *CosmosdbAccountCapacity {
	var returns *CosmosdbAccountCapacity
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) ConnectionStrings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectionStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) ConsistencyPolicy() CosmosdbAccountConsistencyPolicyOutputReference {
	var returns CosmosdbAccountConsistencyPolicyOutputReference
	_jsii_.Get(
		j,
		"consistencyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) ConsistencyPolicyInput() *CosmosdbAccountConsistencyPolicy {
	var returns *CosmosdbAccountConsistencyPolicy
	_jsii_.Get(
		j,
		"consistencyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) CorsRule() CosmosdbAccountCorsRuleOutputReference {
	var returns CosmosdbAccountCorsRuleOutputReference
	_jsii_.Get(
		j,
		"corsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) CorsRuleInput() *CosmosdbAccountCorsRule {
	var returns *CosmosdbAccountCorsRule
	_jsii_.Get(
		j,
		"corsRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) CreateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) CreateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) DefaultIdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIdentityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) DefaultIdentityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIdentityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) EnableAutomaticFailover() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticFailover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) EnableAutomaticFailoverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticFailoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) EnableFreeTier() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFreeTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) EnableFreeTierInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFreeTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) EnableMultipleWriteLocations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultipleWriteLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) EnableMultipleWriteLocationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultipleWriteLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) GeoLocation() CosmosdbAccountGeoLocationList {
	var returns CosmosdbAccountGeoLocationList
	_jsii_.Get(
		j,
		"geoLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) GeoLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Identity() CosmosdbAccountIdentityOutputReference {
	var returns CosmosdbAccountIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) IdentityInput() *CosmosdbAccountIdentity {
	var returns *CosmosdbAccountIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) IpRangeFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipRangeFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) IpRangeFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipRangeFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) IsVirtualNetworkFilterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isVirtualNetworkFilterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) IsVirtualNetworkFilterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isVirtualNetworkFilterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) KeyVaultKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) KeyVaultKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) LocalAuthenticationDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthenticationDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) LocalAuthenticationDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthenticationDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) MongoServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) MongoServerVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoServerVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) NetworkAclBypassForAzureServices() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkAclBypassForAzureServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) NetworkAclBypassForAzureServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkAclBypassForAzureServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) NetworkAclBypassIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkAclBypassIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) NetworkAclBypassIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkAclBypassIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) OfferType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) OfferTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) PrimaryKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) PrimaryReadonlyKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryReadonlyKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) ReadEndpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Restore() CosmosdbAccountRestoreOutputReference {
	var returns CosmosdbAccountRestoreOutputReference
	_jsii_.Get(
		j,
		"restore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) RestoreInput() *CosmosdbAccountRestore {
	var returns *CosmosdbAccountRestore
	_jsii_.Get(
		j,
		"restoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) SecondaryKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) SecondaryReadonlyKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryReadonlyKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) Timeouts() CosmosdbAccountTimeoutsOutputReference {
	var returns CosmosdbAccountTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) VirtualNetworkRule() CosmosdbAccountVirtualNetworkRuleList {
	var returns CosmosdbAccountVirtualNetworkRuleList
	_jsii_.Get(
		j,
		"virtualNetworkRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) VirtualNetworkRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtualNetworkRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) WriteEndpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"writeEndpoints",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account azurerm_cosmosdb_account} Resource.
func NewCosmosdbAccount(scope constructs.Construct, id *string, config *CosmosdbAccountConfig) CosmosdbAccount {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccount{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account azurerm_cosmosdb_account} Resource.
func NewCosmosdbAccount_Override(c CosmosdbAccount, scope constructs.Construct, id *string, config *CosmosdbAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccount",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetAccessKeyMetadataWritesEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"accessKeyMetadataWritesEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetAnalyticalStorageEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"analyticalStorageEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetCreateMode(val *string) {
	_jsii_.Set(
		j,
		"createMode",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetDefaultIdentityType(val *string) {
	_jsii_.Set(
		j,
		"defaultIdentityType",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetEnableAutomaticFailover(val interface{}) {
	_jsii_.Set(
		j,
		"enableAutomaticFailover",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetEnableFreeTier(val interface{}) {
	_jsii_.Set(
		j,
		"enableFreeTier",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetEnableMultipleWriteLocations(val interface{}) {
	_jsii_.Set(
		j,
		"enableMultipleWriteLocations",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetIpRangeFilter(val *string) {
	_jsii_.Set(
		j,
		"ipRangeFilter",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetIsVirtualNetworkFilterEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"isVirtualNetworkFilterEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetKeyVaultKeyId(val *string) {
	_jsii_.Set(
		j,
		"keyVaultKeyId",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetKind(val *string) {
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetLocalAuthenticationDisabled(val interface{}) {
	_jsii_.Set(
		j,
		"localAuthenticationDisabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetMongoServerVersion(val *string) {
	_jsii_.Set(
		j,
		"mongoServerVersion",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetNetworkAclBypassForAzureServices(val interface{}) {
	_jsii_.Set(
		j,
		"networkAclBypassForAzureServices",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetNetworkAclBypassIds(val *[]*string) {
	_jsii_.Set(
		j,
		"networkAclBypassIds",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetOfferType(val *string) {
	_jsii_.Set(
		j,
		"offerType",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetPublicNetworkAccessEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount) SetTags(val *map[string]*string) {
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
func CosmosdbAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CosmosdbAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CosmosdbAccount) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutAnalyticalStorage(value *CosmosdbAccountAnalyticalStorage) {
	_jsii_.InvokeVoid(
		c,
		"putAnalyticalStorage",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutBackup(value *CosmosdbAccountBackup) {
	_jsii_.InvokeVoid(
		c,
		"putBackup",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutCapabilities(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putCapabilities",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutCapacity(value *CosmosdbAccountCapacity) {
	_jsii_.InvokeVoid(
		c,
		"putCapacity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutConsistencyPolicy(value *CosmosdbAccountConsistencyPolicy) {
	_jsii_.InvokeVoid(
		c,
		"putConsistencyPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutCorsRule(value *CosmosdbAccountCorsRule) {
	_jsii_.InvokeVoid(
		c,
		"putCorsRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutGeoLocation(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putGeoLocation",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutIdentity(value *CosmosdbAccountIdentity) {
	_jsii_.InvokeVoid(
		c,
		"putIdentity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutRestore(value *CosmosdbAccountRestore) {
	_jsii_.InvokeVoid(
		c,
		"putRestore",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutTimeouts(value *CosmosdbAccountTimeouts) {
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutVirtualNetworkRule(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putVirtualNetworkRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetAccessKeyMetadataWritesEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessKeyMetadataWritesEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetAnalyticalStorage() {
	_jsii_.InvokeVoid(
		c,
		"resetAnalyticalStorage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetAnalyticalStorageEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAnalyticalStorageEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetBackup() {
	_jsii_.InvokeVoid(
		c,
		"resetBackup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetCapabilities() {
	_jsii_.InvokeVoid(
		c,
		"resetCapabilities",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetCapacity() {
	_jsii_.InvokeVoid(
		c,
		"resetCapacity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetCorsRule() {
	_jsii_.InvokeVoid(
		c,
		"resetCorsRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetCreateMode() {
	_jsii_.InvokeVoid(
		c,
		"resetCreateMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetDefaultIdentityType() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultIdentityType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetEnableAutomaticFailover() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableAutomaticFailover",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetEnableFreeTier() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableFreeTier",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetEnableMultipleWriteLocations() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableMultipleWriteLocations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetIdentity() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetIpRangeFilter() {
	_jsii_.InvokeVoid(
		c,
		"resetIpRangeFilter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetIsVirtualNetworkFilterEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetIsVirtualNetworkFilterEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetKeyVaultKeyId() {
	_jsii_.InvokeVoid(
		c,
		"resetKeyVaultKeyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetKind() {
	_jsii_.InvokeVoid(
		c,
		"resetKind",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetLocalAuthenticationDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalAuthenticationDisabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetMongoServerVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetMongoServerVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetNetworkAclBypassForAzureServices() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkAclBypassForAzureServices",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetNetworkAclBypassIds() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkAclBypassIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetRestore() {
	_jsii_.InvokeVoid(
		c,
		"resetRestore",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetVirtualNetworkRule() {
	_jsii_.InvokeVoid(
		c,
		"resetVirtualNetworkRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountAnalyticalStorage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#schema_type CosmosdbAccount#schema_type}.
	SchemaType *string `field:"required" json:"schemaType" yaml:"schemaType"`
}

type CosmosdbAccountAnalyticalStorageOutputReference interface {
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
	InternalValue() *CosmosdbAccountAnalyticalStorage
	SetInternalValue(val *CosmosdbAccountAnalyticalStorage)
	SchemaType() *string
	SetSchemaType(val *string)
	SchemaTypeInput() *string
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

// The jsii proxy struct for CosmosdbAccountAnalyticalStorageOutputReference
type jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) InternalValue() *CosmosdbAccountAnalyticalStorage {
	var returns *CosmosdbAccountAnalyticalStorage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) SchemaType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) SchemaTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountAnalyticalStorageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CosmosdbAccountAnalyticalStorageOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountAnalyticalStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCosmosdbAccountAnalyticalStorageOutputReference_Override(c CosmosdbAccountAnalyticalStorageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountAnalyticalStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) SetInternalValue(val *CosmosdbAccountAnalyticalStorage) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) SetSchemaType(val *string) {
	_jsii_.Set(
		j,
		"schemaType",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountAnalyticalStorageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountBackup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#type CosmosdbAccount#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#interval_in_minutes CosmosdbAccount#interval_in_minutes}.
	IntervalInMinutes *float64 `field:"optional" json:"intervalInMinutes" yaml:"intervalInMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#retention_in_hours CosmosdbAccount#retention_in_hours}.
	RetentionInHours *float64 `field:"optional" json:"retentionInHours" yaml:"retentionInHours"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#storage_redundancy CosmosdbAccount#storage_redundancy}.
	StorageRedundancy *string `field:"optional" json:"storageRedundancy" yaml:"storageRedundancy"`
}

type CosmosdbAccountBackupOutputReference interface {
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
	InternalValue() *CosmosdbAccountBackup
	SetInternalValue(val *CosmosdbAccountBackup)
	IntervalInMinutes() *float64
	SetIntervalInMinutes(val *float64)
	IntervalInMinutesInput() *float64
	RetentionInHours() *float64
	SetRetentionInHours(val *float64)
	RetentionInHoursInput() *float64
	StorageRedundancy() *string
	SetStorageRedundancy(val *string)
	StorageRedundancyInput() *string
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
	ResetIntervalInMinutes()
	ResetRetentionInHours()
	ResetStorageRedundancy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CosmosdbAccountBackupOutputReference
type jsiiProxy_CosmosdbAccountBackupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) InternalValue() *CosmosdbAccountBackup {
	var returns *CosmosdbAccountBackup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) IntervalInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) IntervalInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) RetentionInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) RetentionInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) StorageRedundancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRedundancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) StorageRedundancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRedundancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountBackupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CosmosdbAccountBackupOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountBackupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCosmosdbAccountBackupOutputReference_Override(c CosmosdbAccountBackupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) SetInternalValue(val *CosmosdbAccountBackup) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) SetIntervalInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"intervalInMinutes",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) SetRetentionInHours(val *float64) {
	_jsii_.Set(
		j,
		"retentionInHours",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) SetStorageRedundancy(val *string) {
	_jsii_.Set(
		j,
		"storageRedundancy",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountBackupOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) ResetIntervalInMinutes() {
	_jsii_.InvokeVoid(
		c,
		"resetIntervalInMinutes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) ResetRetentionInHours() {
	_jsii_.InvokeVoid(
		c,
		"resetRetentionInHours",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) ResetStorageRedundancy() {
	_jsii_.InvokeVoid(
		c,
		"resetStorageRedundancy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountBackupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountCapabilities struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#name CosmosdbAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

type CosmosdbAccountCapabilitiesList interface {
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
	Get(index *float64) CosmosdbAccountCapabilitiesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CosmosdbAccountCapabilitiesList
type jsiiProxy_CosmosdbAccountCapabilitiesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountCapabilitiesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CosmosdbAccountCapabilitiesList {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountCapabilitiesList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountCapabilitiesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCosmosdbAccountCapabilitiesList_Override(c CosmosdbAccountCapabilitiesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountCapabilitiesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesList) Get(index *float64) CosmosdbAccountCapabilitiesOutputReference {
	var returns CosmosdbAccountCapabilitiesOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountCapabilitiesOutputReference interface {
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

// The jsii proxy struct for CosmosdbAccountCapabilitiesOutputReference
type jsiiProxy_CosmosdbAccountCapabilitiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountCapabilitiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CosmosdbAccountCapabilitiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountCapabilitiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCosmosdbAccountCapabilitiesOutputReference_Override(c CosmosdbAccountCapabilitiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapabilitiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountCapacity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#total_throughput_limit CosmosdbAccount#total_throughput_limit}.
	TotalThroughputLimit *float64 `field:"required" json:"totalThroughputLimit" yaml:"totalThroughputLimit"`
}

type CosmosdbAccountCapacityOutputReference interface {
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
	InternalValue() *CosmosdbAccountCapacity
	SetInternalValue(val *CosmosdbAccountCapacity)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalThroughputLimit() *float64
	SetTotalThroughputLimit(val *float64)
	TotalThroughputLimitInput() *float64
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

// The jsii proxy struct for CosmosdbAccountCapacityOutputReference
type jsiiProxy_CosmosdbAccountCapacityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbAccountCapacityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapacityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapacityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapacityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapacityOutputReference) InternalValue() *CosmosdbAccountCapacity {
	var returns *CosmosdbAccountCapacity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapacityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapacityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapacityOutputReference) TotalThroughputLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalThroughputLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCapacityOutputReference) TotalThroughputLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalThroughputLimitInput",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountCapacityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CosmosdbAccountCapacityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountCapacityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountCapacityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCosmosdbAccountCapacityOutputReference_Override(c CosmosdbAccountCapacityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountCapacityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapacityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapacityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapacityOutputReference) SetInternalValue(val *CosmosdbAccountCapacity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapacityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapacityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCapacityOutputReference) SetTotalThroughputLimit(val *float64) {
	_jsii_.Set(
		j,
		"totalThroughputLimit",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountCapacityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapacityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapacityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapacityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapacityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapacityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapacityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapacityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapacityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapacityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapacityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapacityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapacityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCapacityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountConfig struct {
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
	// consistency_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#consistency_policy CosmosdbAccount#consistency_policy}
	ConsistencyPolicy *CosmosdbAccountConsistencyPolicy `field:"required" json:"consistencyPolicy" yaml:"consistencyPolicy"`
	// geo_location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#geo_location CosmosdbAccount#geo_location}
	GeoLocation interface{} `field:"required" json:"geoLocation" yaml:"geoLocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#location CosmosdbAccount#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#name CosmosdbAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#offer_type CosmosdbAccount#offer_type}.
	OfferType *string `field:"required" json:"offerType" yaml:"offerType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#resource_group_name CosmosdbAccount#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#access_key_metadata_writes_enabled CosmosdbAccount#access_key_metadata_writes_enabled}.
	AccessKeyMetadataWritesEnabled interface{} `field:"optional" json:"accessKeyMetadataWritesEnabled" yaml:"accessKeyMetadataWritesEnabled"`
	// analytical_storage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#analytical_storage CosmosdbAccount#analytical_storage}
	AnalyticalStorage *CosmosdbAccountAnalyticalStorage `field:"optional" json:"analyticalStorage" yaml:"analyticalStorage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#analytical_storage_enabled CosmosdbAccount#analytical_storage_enabled}.
	AnalyticalStorageEnabled interface{} `field:"optional" json:"analyticalStorageEnabled" yaml:"analyticalStorageEnabled"`
	// backup block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#backup CosmosdbAccount#backup}
	Backup *CosmosdbAccountBackup `field:"optional" json:"backup" yaml:"backup"`
	// capabilities block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#capabilities CosmosdbAccount#capabilities}
	Capabilities interface{} `field:"optional" json:"capabilities" yaml:"capabilities"`
	// capacity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#capacity CosmosdbAccount#capacity}
	Capacity *CosmosdbAccountCapacity `field:"optional" json:"capacity" yaml:"capacity"`
	// cors_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#cors_rule CosmosdbAccount#cors_rule}
	CorsRule *CosmosdbAccountCorsRule `field:"optional" json:"corsRule" yaml:"corsRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#create_mode CosmosdbAccount#create_mode}.
	CreateMode *string `field:"optional" json:"createMode" yaml:"createMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#default_identity_type CosmosdbAccount#default_identity_type}.
	DefaultIdentityType *string `field:"optional" json:"defaultIdentityType" yaml:"defaultIdentityType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#enable_automatic_failover CosmosdbAccount#enable_automatic_failover}.
	EnableAutomaticFailover interface{} `field:"optional" json:"enableAutomaticFailover" yaml:"enableAutomaticFailover"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#enable_free_tier CosmosdbAccount#enable_free_tier}.
	EnableFreeTier interface{} `field:"optional" json:"enableFreeTier" yaml:"enableFreeTier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#enable_multiple_write_locations CosmosdbAccount#enable_multiple_write_locations}.
	EnableMultipleWriteLocations interface{} `field:"optional" json:"enableMultipleWriteLocations" yaml:"enableMultipleWriteLocations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#id CosmosdbAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#identity CosmosdbAccount#identity}
	Identity *CosmosdbAccountIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#ip_range_filter CosmosdbAccount#ip_range_filter}.
	IpRangeFilter *string `field:"optional" json:"ipRangeFilter" yaml:"ipRangeFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#is_virtual_network_filter_enabled CosmosdbAccount#is_virtual_network_filter_enabled}.
	IsVirtualNetworkFilterEnabled interface{} `field:"optional" json:"isVirtualNetworkFilterEnabled" yaml:"isVirtualNetworkFilterEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#key_vault_key_id CosmosdbAccount#key_vault_key_id}.
	KeyVaultKeyId *string `field:"optional" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#kind CosmosdbAccount#kind}.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#local_authentication_disabled CosmosdbAccount#local_authentication_disabled}.
	LocalAuthenticationDisabled interface{} `field:"optional" json:"localAuthenticationDisabled" yaml:"localAuthenticationDisabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#mongo_server_version CosmosdbAccount#mongo_server_version}.
	MongoServerVersion *string `field:"optional" json:"mongoServerVersion" yaml:"mongoServerVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#network_acl_bypass_for_azure_services CosmosdbAccount#network_acl_bypass_for_azure_services}.
	NetworkAclBypassForAzureServices interface{} `field:"optional" json:"networkAclBypassForAzureServices" yaml:"networkAclBypassForAzureServices"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#network_acl_bypass_ids CosmosdbAccount#network_acl_bypass_ids}.
	NetworkAclBypassIds *[]*string `field:"optional" json:"networkAclBypassIds" yaml:"networkAclBypassIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#public_network_access_enabled CosmosdbAccount#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// restore block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#restore CosmosdbAccount#restore}
	Restore *CosmosdbAccountRestore `field:"optional" json:"restore" yaml:"restore"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#tags CosmosdbAccount#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#timeouts CosmosdbAccount#timeouts}
	Timeouts *CosmosdbAccountTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// virtual_network_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#virtual_network_rule CosmosdbAccount#virtual_network_rule}
	VirtualNetworkRule interface{} `field:"optional" json:"virtualNetworkRule" yaml:"virtualNetworkRule"`
}

type CosmosdbAccountConsistencyPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#consistency_level CosmosdbAccount#consistency_level}.
	ConsistencyLevel *string `field:"required" json:"consistencyLevel" yaml:"consistencyLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#max_interval_in_seconds CosmosdbAccount#max_interval_in_seconds}.
	MaxIntervalInSeconds *float64 `field:"optional" json:"maxIntervalInSeconds" yaml:"maxIntervalInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#max_staleness_prefix CosmosdbAccount#max_staleness_prefix}.
	MaxStalenessPrefix *float64 `field:"optional" json:"maxStalenessPrefix" yaml:"maxStalenessPrefix"`
}

type CosmosdbAccountConsistencyPolicyOutputReference interface {
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
	ConsistencyLevel() *string
	SetConsistencyLevel(val *string)
	ConsistencyLevelInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CosmosdbAccountConsistencyPolicy
	SetInternalValue(val *CosmosdbAccountConsistencyPolicy)
	MaxIntervalInSeconds() *float64
	SetMaxIntervalInSeconds(val *float64)
	MaxIntervalInSecondsInput() *float64
	MaxStalenessPrefix() *float64
	SetMaxStalenessPrefix(val *float64)
	MaxStalenessPrefixInput() *float64
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
	ResetMaxIntervalInSeconds()
	ResetMaxStalenessPrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CosmosdbAccountConsistencyPolicyOutputReference
type jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) ConsistencyLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consistencyLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) ConsistencyLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consistencyLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) InternalValue() *CosmosdbAccountConsistencyPolicy {
	var returns *CosmosdbAccountConsistencyPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) MaxIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) MaxIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) MaxStalenessPrefix() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStalenessPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) MaxStalenessPrefixInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStalenessPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountConsistencyPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CosmosdbAccountConsistencyPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountConsistencyPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCosmosdbAccountConsistencyPolicyOutputReference_Override(c CosmosdbAccountConsistencyPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountConsistencyPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) SetConsistencyLevel(val *string) {
	_jsii_.Set(
		j,
		"consistencyLevel",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) SetInternalValue(val *CosmosdbAccountConsistencyPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) SetMaxIntervalInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maxIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) SetMaxStalenessPrefix(val *float64) {
	_jsii_.Set(
		j,
		"maxStalenessPrefix",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) ResetMaxIntervalInSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxIntervalInSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) ResetMaxStalenessPrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxStalenessPrefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountConsistencyPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountCorsRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#allowed_headers CosmosdbAccount#allowed_headers}.
	AllowedHeaders *[]*string `field:"required" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#allowed_methods CosmosdbAccount#allowed_methods}.
	AllowedMethods *[]*string `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#allowed_origins CosmosdbAccount#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#exposed_headers CosmosdbAccount#exposed_headers}.
	ExposedHeaders *[]*string `field:"required" json:"exposedHeaders" yaml:"exposedHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#max_age_in_seconds CosmosdbAccount#max_age_in_seconds}.
	MaxAgeInSeconds *float64 `field:"required" json:"maxAgeInSeconds" yaml:"maxAgeInSeconds"`
}

type CosmosdbAccountCorsRuleOutputReference interface {
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
	InternalValue() *CosmosdbAccountCorsRule
	SetInternalValue(val *CosmosdbAccountCorsRule)
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

// The jsii proxy struct for CosmosdbAccountCorsRuleOutputReference
type jsiiProxy_CosmosdbAccountCorsRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) AllowedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) AllowedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) AllowedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) AllowedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) AllowedOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) AllowedOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) ExposedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) ExposedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) InternalValue() *CosmosdbAccountCorsRule {
	var returns *CosmosdbAccountCorsRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) MaxAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) MaxAgeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountCorsRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CosmosdbAccountCorsRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountCorsRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountCorsRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCosmosdbAccountCorsRuleOutputReference_Override(c CosmosdbAccountCorsRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountCorsRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) SetAllowedHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedHeaders",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) SetAllowedMethods(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedMethods",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) SetAllowedOrigins(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedOrigins",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) SetExposedHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"exposedHeaders",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) SetInternalValue(val *CosmosdbAccountCorsRule) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) SetMaxAgeInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maxAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountCorsRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountGeoLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#failover_priority CosmosdbAccount#failover_priority}.
	FailoverPriority *float64 `field:"required" json:"failoverPriority" yaml:"failoverPriority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#location CosmosdbAccount#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#zone_redundant CosmosdbAccount#zone_redundant}.
	ZoneRedundant interface{} `field:"optional" json:"zoneRedundant" yaml:"zoneRedundant"`
}

type CosmosdbAccountGeoLocationList interface {
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
	Get(index *float64) CosmosdbAccountGeoLocationOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CosmosdbAccountGeoLocationList
type jsiiProxy_CosmosdbAccountGeoLocationList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountGeoLocationList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CosmosdbAccountGeoLocationList {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountGeoLocationList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountGeoLocationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCosmosdbAccountGeoLocationList_Override(c CosmosdbAccountGeoLocationList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountGeoLocationList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationList) Get(index *float64) CosmosdbAccountGeoLocationOutputReference {
	var returns CosmosdbAccountGeoLocationOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountGeoLocationOutputReference interface {
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
	FailoverPriority() *float64
	SetFailoverPriority(val *float64)
	FailoverPriorityInput() *float64
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ZoneRedundant() interface{}
	SetZoneRedundant(val interface{})
	ZoneRedundantInput() interface{}
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
	ResetZoneRedundant()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CosmosdbAccountGeoLocationOutputReference
type jsiiProxy_CosmosdbAccountGeoLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) FailoverPriority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failoverPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) FailoverPriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failoverPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) ZoneRedundant() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneRedundant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) ZoneRedundantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneRedundantInput",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountGeoLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CosmosdbAccountGeoLocationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountGeoLocationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountGeoLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCosmosdbAccountGeoLocationOutputReference_Override(c CosmosdbAccountGeoLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountGeoLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) SetFailoverPriority(val *float64) {
	_jsii_.Set(
		j,
		"failoverPriority",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) SetZoneRedundant(val interface{}) {
	_jsii_.Set(
		j,
		"zoneRedundant",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) ResetZoneRedundant() {
	_jsii_.InvokeVoid(
		c,
		"resetZoneRedundant",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountGeoLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#type CosmosdbAccount#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#identity_ids CosmosdbAccount#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

type CosmosdbAccountIdentityOutputReference interface {
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
	InternalValue() *CosmosdbAccountIdentity
	SetInternalValue(val *CosmosdbAccountIdentity)
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

// The jsii proxy struct for CosmosdbAccountIdentityOutputReference
type jsiiProxy_CosmosdbAccountIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) IdentityIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) IdentityIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) InternalValue() *CosmosdbAccountIdentity {
	var returns *CosmosdbAccountIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CosmosdbAccountIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCosmosdbAccountIdentityOutputReference_Override(c CosmosdbAccountIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) SetIdentityIds(val *[]*string) {
	_jsii_.Set(
		j,
		"identityIds",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) SetInternalValue(val *CosmosdbAccountIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountIdentityOutputReference) ResetIdentityIds() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentityIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountRestore struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#restore_timestamp_in_utc CosmosdbAccount#restore_timestamp_in_utc}.
	RestoreTimestampInUtc *string `field:"required" json:"restoreTimestampInUtc" yaml:"restoreTimestampInUtc"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#source_cosmosdb_account_id CosmosdbAccount#source_cosmosdb_account_id}.
	SourceCosmosdbAccountId *string `field:"required" json:"sourceCosmosdbAccountId" yaml:"sourceCosmosdbAccountId"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#database CosmosdbAccount#database}
	Database interface{} `field:"optional" json:"database" yaml:"database"`
}

type CosmosdbAccountRestoreDatabase struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#name CosmosdbAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#collection_names CosmosdbAccount#collection_names}.
	CollectionNames *[]*string `field:"optional" json:"collectionNames" yaml:"collectionNames"`
}

type CosmosdbAccountRestoreDatabaseList interface {
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
	Get(index *float64) CosmosdbAccountRestoreDatabaseOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CosmosdbAccountRestoreDatabaseList
type jsiiProxy_CosmosdbAccountRestoreDatabaseList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountRestoreDatabaseList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CosmosdbAccountRestoreDatabaseList {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountRestoreDatabaseList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountRestoreDatabaseList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCosmosdbAccountRestoreDatabaseList_Override(c CosmosdbAccountRestoreDatabaseList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountRestoreDatabaseList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseList) Get(index *float64) CosmosdbAccountRestoreDatabaseOutputReference {
	var returns CosmosdbAccountRestoreDatabaseOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountRestoreDatabaseOutputReference interface {
	cdktf.ComplexObject
	CollectionNames() *[]*string
	SetCollectionNames(val *[]*string)
	CollectionNamesInput() *[]*string
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
	ResetCollectionNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CosmosdbAccountRestoreDatabaseOutputReference
type jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) CollectionNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"collectionNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) CollectionNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"collectionNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountRestoreDatabaseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CosmosdbAccountRestoreDatabaseOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountRestoreDatabaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCosmosdbAccountRestoreDatabaseOutputReference_Override(c CosmosdbAccountRestoreDatabaseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountRestoreDatabaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) SetCollectionNames(val *[]*string) {
	_jsii_.Set(
		j,
		"collectionNames",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) ResetCollectionNames() {
	_jsii_.InvokeVoid(
		c,
		"resetCollectionNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreDatabaseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountRestoreOutputReference interface {
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
	Database() CosmosdbAccountRestoreDatabaseList
	DatabaseInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *CosmosdbAccountRestore
	SetInternalValue(val *CosmosdbAccountRestore)
	RestoreTimestampInUtc() *string
	SetRestoreTimestampInUtc(val *string)
	RestoreTimestampInUtcInput() *string
	SourceCosmosdbAccountId() *string
	SetSourceCosmosdbAccountId(val *string)
	SourceCosmosdbAccountIdInput() *string
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
	PutDatabase(value interface{})
	ResetDatabase()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CosmosdbAccountRestoreOutputReference
type jsiiProxy_CosmosdbAccountRestoreOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) Database() CosmosdbAccountRestoreDatabaseList {
	var returns CosmosdbAccountRestoreDatabaseList
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) InternalValue() *CosmosdbAccountRestore {
	var returns *CosmosdbAccountRestore
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) RestoreTimestampInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTimestampInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) RestoreTimestampInUtcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTimestampInUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) SourceCosmosdbAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCosmosdbAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) SourceCosmosdbAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCosmosdbAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountRestoreOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CosmosdbAccountRestoreOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountRestoreOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountRestoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCosmosdbAccountRestoreOutputReference_Override(c CosmosdbAccountRestoreOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountRestoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) SetInternalValue(val *CosmosdbAccountRestore) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) SetRestoreTimestampInUtc(val *string) {
	_jsii_.Set(
		j,
		"restoreTimestampInUtc",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) SetSourceCosmosdbAccountId(val *string) {
	_jsii_.Set(
		j,
		"sourceCosmosdbAccountId",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountRestoreOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) PutDatabase(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putDatabase",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		c,
		"resetDatabase",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountRestoreOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#create CosmosdbAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#delete CosmosdbAccount#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#read CosmosdbAccount#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#update CosmosdbAccount#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type CosmosdbAccountTimeoutsOutputReference interface {
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

// The jsii proxy struct for CosmosdbAccountTimeoutsOutputReference
type jsiiProxy_CosmosdbAccountTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CosmosdbAccountTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCosmosdbAccountTimeoutsOutputReference_Override(c CosmosdbAccountTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		c,
		"resetCreate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		c,
		"resetDelete",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		c,
		"resetRead",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountVirtualNetworkRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#id CosmosdbAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cosmosdb_account#ignore_missing_vnet_service_endpoint CosmosdbAccount#ignore_missing_vnet_service_endpoint}.
	IgnoreMissingVnetServiceEndpoint interface{} `field:"optional" json:"ignoreMissingVnetServiceEndpoint" yaml:"ignoreMissingVnetServiceEndpoint"`
}

type CosmosdbAccountVirtualNetworkRuleList interface {
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
	Get(index *float64) CosmosdbAccountVirtualNetworkRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CosmosdbAccountVirtualNetworkRuleList
type jsiiProxy_CosmosdbAccountVirtualNetworkRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountVirtualNetworkRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CosmosdbAccountVirtualNetworkRuleList {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountVirtualNetworkRuleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountVirtualNetworkRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCosmosdbAccountVirtualNetworkRuleList_Override(c CosmosdbAccountVirtualNetworkRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountVirtualNetworkRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleList) Get(index *float64) CosmosdbAccountVirtualNetworkRuleOutputReference {
	var returns CosmosdbAccountVirtualNetworkRuleOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CosmosdbAccountVirtualNetworkRuleOutputReference interface {
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	IgnoreMissingVnetServiceEndpoint() interface{}
	SetIgnoreMissingVnetServiceEndpoint(val interface{})
	IgnoreMissingVnetServiceEndpointInput() interface{}
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
	ResetIgnoreMissingVnetServiceEndpoint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CosmosdbAccountVirtualNetworkRuleOutputReference
type jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) IgnoreMissingVnetServiceEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMissingVnetServiceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) IgnoreMissingVnetServiceEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMissingVnetServiceEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCosmosdbAccountVirtualNetworkRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CosmosdbAccountVirtualNetworkRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountVirtualNetworkRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCosmosdbAccountVirtualNetworkRuleOutputReference_Override(c CosmosdbAccountVirtualNetworkRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccountVirtualNetworkRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) SetIgnoreMissingVnetServiceEndpoint(val interface{}) {
	_jsii_.Set(
		j,
		"ignoreMissingVnetServiceEndpoint",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) ResetIgnoreMissingVnetServiceEndpoint() {
	_jsii_.InvokeVoid(
		c,
		"resetIgnoreMissingVnetServiceEndpoint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccountVirtualNetworkRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

