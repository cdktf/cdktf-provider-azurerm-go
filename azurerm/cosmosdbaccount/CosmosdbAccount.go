// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/cosmosdbaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/cosmosdb_account azurerm_cosmosdb_account}.
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
	AutomaticFailoverEnabled() interface{}
	SetAutomaticFailoverEnabled(val interface{})
	AutomaticFailoverEnabledInput() interface{}
	Backup() CosmosdbAccountBackupOutputReference
	BackupInput() *CosmosdbAccountBackup
	BurstCapacityEnabled() interface{}
	SetBurstCapacityEnabled(val interface{})
	BurstCapacityEnabledInput() interface{}
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
	ConsistencyPolicy() CosmosdbAccountConsistencyPolicyOutputReference
	ConsistencyPolicyInput() *CosmosdbAccountConsistencyPolicy
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CorsRule() CosmosdbAccountCorsRuleOutputReference
	CorsRuleInput() *CosmosdbAccountCorsRule
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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
	Endpoint() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	FreeTierEnabled() interface{}
	SetFreeTierEnabled(val interface{})
	FreeTierEnabledInput() interface{}
	// Experimental.
	FriendlyUniqueId() *string
	GeoLocation() CosmosdbAccountGeoLocationList
	GeoLocationInput() interface{}
	Id() *string
	SetId(val *string)
	Identity() CosmosdbAccountIdentityOutputReference
	IdentityInput() *CosmosdbAccountIdentity
	IdInput() *string
	IpRangeFilter() *[]*string
	SetIpRangeFilter(val *[]*string)
	IpRangeFilterInput() *[]*string
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
	ManagedHsmKeyId() *string
	SetManagedHsmKeyId(val *string)
	ManagedHsmKeyIdInput() *string
	MinimalTlsVersion() *string
	SetMinimalTlsVersion(val *string)
	MinimalTlsVersionInput() *string
	MongoServerVersion() *string
	SetMongoServerVersion(val *string)
	MongoServerVersionInput() *string
	MultipleWriteLocationsEnabled() interface{}
	SetMultipleWriteLocationsEnabled(val interface{})
	MultipleWriteLocationsEnabledInput() interface{}
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
	PartitionMergeEnabled() interface{}
	SetPartitionMergeEnabled(val interface{})
	PartitionMergeEnabledInput() interface{}
	PrimaryKey() *string
	PrimaryMongodbConnectionString() *string
	PrimaryReadonlyKey() *string
	PrimaryReadonlyMongodbConnectionString() *string
	PrimaryReadonlySqlConnectionString() *string
	PrimarySqlConnectionString() *string
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
	SecondaryMongodbConnectionString() *string
	SecondaryReadonlyKey() *string
	SecondaryReadonlyMongodbConnectionString() *string
	SecondaryReadonlySqlConnectionString() *string
	SecondarySqlConnectionString() *string
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
	ResetAutomaticFailoverEnabled()
	ResetBackup()
	ResetBurstCapacityEnabled()
	ResetCapabilities()
	ResetCapacity()
	ResetCorsRule()
	ResetCreateMode()
	ResetDefaultIdentityType()
	ResetFreeTierEnabled()
	ResetId()
	ResetIdentity()
	ResetIpRangeFilter()
	ResetIsVirtualNetworkFilterEnabled()
	ResetKeyVaultKeyId()
	ResetKind()
	ResetLocalAuthenticationDisabled()
	ResetManagedHsmKeyId()
	ResetMinimalTlsVersion()
	ResetMongoServerVersion()
	ResetMultipleWriteLocationsEnabled()
	ResetNetworkAclBypassForAzureServices()
	ResetNetworkAclBypassIds()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartitionMergeEnabled()
	ResetPublicNetworkAccessEnabled()
	ResetRestore()
	ResetTags()
	ResetTimeouts()
	ResetVirtualNetworkRule()
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

func (j *jsiiProxy_CosmosdbAccount) AutomaticFailoverEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticFailoverEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) AutomaticFailoverEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticFailoverEnabledInput",
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

func (j *jsiiProxy_CosmosdbAccount) BurstCapacityEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"burstCapacityEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) BurstCapacityEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"burstCapacityEnabledInput",
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

func (j *jsiiProxy_CosmosdbAccount) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_CosmosdbAccount) FreeTierEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"freeTierEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) FreeTierEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"freeTierEnabledInput",
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

func (j *jsiiProxy_CosmosdbAccount) IpRangeFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipRangeFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) IpRangeFilterInput() *[]*string {
	var returns *[]*string
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

func (j *jsiiProxy_CosmosdbAccount) ManagedHsmKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedHsmKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) ManagedHsmKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedHsmKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) MinimalTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimalTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) MinimalTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimalTlsVersionInput",
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

func (j *jsiiProxy_CosmosdbAccount) MultipleWriteLocationsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multipleWriteLocationsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) MultipleWriteLocationsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multipleWriteLocationsEnabledInput",
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

func (j *jsiiProxy_CosmosdbAccount) PartitionMergeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionMergeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) PartitionMergeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionMergeEnabledInput",
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

func (j *jsiiProxy_CosmosdbAccount) PrimaryMongodbConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryMongodbConnectionString",
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

func (j *jsiiProxy_CosmosdbAccount) PrimaryReadonlyMongodbConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryReadonlyMongodbConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) PrimaryReadonlySqlConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryReadonlySqlConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) PrimarySqlConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primarySqlConnectionString",
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

func (j *jsiiProxy_CosmosdbAccount) SecondaryMongodbConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryMongodbConnectionString",
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

func (j *jsiiProxy_CosmosdbAccount) SecondaryReadonlyMongodbConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryReadonlyMongodbConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) SecondaryReadonlySqlConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryReadonlySqlConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbAccount) SecondarySqlConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondarySqlConnectionString",
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/cosmosdb_account azurerm_cosmosdb_account} Resource.
func NewCosmosdbAccount(scope constructs.Construct, id *string, config *CosmosdbAccountConfig) CosmosdbAccount {
	_init_.Initialize()

	if err := validateNewCosmosdbAccountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CosmosdbAccount{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/cosmosdb_account azurerm_cosmosdb_account} Resource.
func NewCosmosdbAccount_Override(c CosmosdbAccount, scope constructs.Construct, id *string, config *CosmosdbAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccount",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetAccessKeyMetadataWritesEnabled(val interface{}) {
	if err := j.validateSetAccessKeyMetadataWritesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessKeyMetadataWritesEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetAnalyticalStorageEnabled(val interface{}) {
	if err := j.validateSetAnalyticalStorageEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analyticalStorageEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetAutomaticFailoverEnabled(val interface{}) {
	if err := j.validateSetAutomaticFailoverEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticFailoverEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetBurstCapacityEnabled(val interface{}) {
	if err := j.validateSetBurstCapacityEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"burstCapacityEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetCreateMode(val *string) {
	if err := j.validateSetCreateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createMode",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetDefaultIdentityType(val *string) {
	if err := j.validateSetDefaultIdentityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultIdentityType",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetFreeTierEnabled(val interface{}) {
	if err := j.validateSetFreeTierEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"freeTierEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetIpRangeFilter(val *[]*string) {
	if err := j.validateSetIpRangeFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipRangeFilter",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetIsVirtualNetworkFilterEnabled(val interface{}) {
	if err := j.validateSetIsVirtualNetworkFilterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isVirtualNetworkFilterEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetKeyVaultKeyId(val *string) {
	if err := j.validateSetKeyVaultKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultKeyId",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetLocalAuthenticationDisabled(val interface{}) {
	if err := j.validateSetLocalAuthenticationDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAuthenticationDisabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetManagedHsmKeyId(val *string) {
	if err := j.validateSetManagedHsmKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedHsmKeyId",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetMinimalTlsVersion(val *string) {
	if err := j.validateSetMinimalTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimalTlsVersion",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetMongoServerVersion(val *string) {
	if err := j.validateSetMongoServerVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mongoServerVersion",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetMultipleWriteLocationsEnabled(val interface{}) {
	if err := j.validateSetMultipleWriteLocationsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multipleWriteLocationsEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetNetworkAclBypassForAzureServices(val interface{}) {
	if err := j.validateSetNetworkAclBypassForAzureServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkAclBypassForAzureServices",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetNetworkAclBypassIds(val *[]*string) {
	if err := j.validateSetNetworkAclBypassIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkAclBypassIds",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetOfferType(val *string) {
	if err := j.validateSetOfferTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offerType",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetPartitionMergeEnabled(val interface{}) {
	if err := j.validateSetPartitionMergeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionMergeEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_CosmosdbAccount)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a CosmosdbAccount resource upon running "cdktf plan <stack-name>".
func CosmosdbAccount_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCosmosdbAccount_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccount",
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
func CosmosdbAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCosmosdbAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CosmosdbAccount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCosmosdbAccount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CosmosdbAccount_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCosmosdbAccount_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.cosmosdbAccount.CosmosdbAccount",
		"isTerraformResource",
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

func (c *jsiiProxy_CosmosdbAccount) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CosmosdbAccount) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CosmosdbAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CosmosdbAccount) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CosmosdbAccount) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CosmosdbAccount) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutAnalyticalStorage(value *CosmosdbAccountAnalyticalStorage) {
	if err := c.validatePutAnalyticalStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAnalyticalStorage",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutBackup(value *CosmosdbAccountBackup) {
	if err := c.validatePutBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBackup",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutCapabilities(value interface{}) {
	if err := c.validatePutCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCapabilities",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutCapacity(value *CosmosdbAccountCapacity) {
	if err := c.validatePutCapacityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCapacity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutConsistencyPolicy(value *CosmosdbAccountConsistencyPolicy) {
	if err := c.validatePutConsistencyPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConsistencyPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutCorsRule(value *CosmosdbAccountCorsRule) {
	if err := c.validatePutCorsRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCorsRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutGeoLocation(value interface{}) {
	if err := c.validatePutGeoLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGeoLocation",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutIdentity(value *CosmosdbAccountIdentity) {
	if err := c.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIdentity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutRestore(value *CosmosdbAccountRestore) {
	if err := c.validatePutRestoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRestore",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutTimeouts(value *CosmosdbAccountTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbAccount) PutVirtualNetworkRule(value interface{}) {
	if err := c.validatePutVirtualNetworkRuleParameters(value); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_CosmosdbAccount) ResetAutomaticFailoverEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAutomaticFailoverEnabled",
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

func (c *jsiiProxy_CosmosdbAccount) ResetBurstCapacityEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetBurstCapacityEnabled",
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

func (c *jsiiProxy_CosmosdbAccount) ResetFreeTierEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetFreeTierEnabled",
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

func (c *jsiiProxy_CosmosdbAccount) ResetManagedHsmKeyId() {
	_jsii_.InvokeVoid(
		c,
		"resetManagedHsmKeyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbAccount) ResetMinimalTlsVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimalTlsVersion",
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

func (c *jsiiProxy_CosmosdbAccount) ResetMultipleWriteLocationsEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetMultipleWriteLocationsEnabled",
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

func (c *jsiiProxy_CosmosdbAccount) ResetPartitionMergeEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetPartitionMergeEnabled",
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

func (c *jsiiProxy_CosmosdbAccount) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbAccount) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
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

