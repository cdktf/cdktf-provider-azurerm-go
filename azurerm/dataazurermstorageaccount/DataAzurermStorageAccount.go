// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermstorageaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/dataazurermstorageaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/data-sources/storage_account azurerm_storage_account}.
type DataAzurermStorageAccount interface {
	cdktf.TerraformDataSource
	AccessTier() *string
	AccountKind() *string
	AccountReplicationType() *string
	AccountTier() *string
	AllowNestedItemsToBePublic() cdktf.IResolvable
	AzureFilesAuthentication() DataAzurermStorageAccountAzureFilesAuthenticationList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomDomain() DataAzurermStorageAccountCustomDomainList
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DnsEndpointType() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpsTrafficOnlyEnabled() cdktf.IResolvable
	Id() *string
	SetId(val *string)
	Identity() DataAzurermStorageAccountIdentityList
	IdInput() *string
	InfrastructureEncryptionEnabled() cdktf.IResolvable
	IsHnsEnabled() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	MinTlsVersion() *string
	SetMinTlsVersion(val *string)
	MinTlsVersionInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfsv3Enabled() cdktf.IResolvable
	// The tree node.
	Node() constructs.Node
	PrimaryAccessKey() *string
	PrimaryBlobConnectionString() *string
	PrimaryBlobEndpoint() *string
	PrimaryBlobHost() *string
	PrimaryBlobInternetEndpoint() *string
	PrimaryBlobInternetHost() *string
	PrimaryBlobMicrosoftEndpoint() *string
	PrimaryBlobMicrosoftHost() *string
	PrimaryConnectionString() *string
	PrimaryDfsEndpoint() *string
	PrimaryDfsHost() *string
	PrimaryDfsInternetEndpoint() *string
	PrimaryDfsInternetHost() *string
	PrimaryDfsMicrosoftEndpoint() *string
	PrimaryDfsMicrosoftHost() *string
	PrimaryFileEndpoint() *string
	PrimaryFileHost() *string
	PrimaryFileInternetEndpoint() *string
	PrimaryFileInternetHost() *string
	PrimaryFileMicrosoftEndpoint() *string
	PrimaryFileMicrosoftHost() *string
	PrimaryLocation() *string
	PrimaryQueueEndpoint() *string
	PrimaryQueueHost() *string
	PrimaryQueueMicrosoftEndpoint() *string
	PrimaryQueueMicrosoftHost() *string
	PrimaryTableEndpoint() *string
	PrimaryTableHost() *string
	PrimaryTableMicrosoftEndpoint() *string
	PrimaryTableMicrosoftHost() *string
	PrimaryWebEndpoint() *string
	PrimaryWebHost() *string
	PrimaryWebInternetEndpoint() *string
	PrimaryWebInternetHost() *string
	PrimaryWebMicrosoftEndpoint() *string
	PrimaryWebMicrosoftHost() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	QueueEncryptionKeyType() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SecondaryAccessKey() *string
	SecondaryBlobConnectionString() *string
	SecondaryBlobEndpoint() *string
	SecondaryBlobHost() *string
	SecondaryBlobInternetEndpoint() *string
	SecondaryBlobInternetHost() *string
	SecondaryBlobMicrosoftEndpoint() *string
	SecondaryBlobMicrosoftHost() *string
	SecondaryConnectionString() *string
	SecondaryDfsEndpoint() *string
	SecondaryDfsHost() *string
	SecondaryDfsInternetEndpoint() *string
	SecondaryDfsInternetHost() *string
	SecondaryDfsMicrosoftEndpoint() *string
	SecondaryDfsMicrosoftHost() *string
	SecondaryFileEndpoint() *string
	SecondaryFileHost() *string
	SecondaryFileInternetEndpoint() *string
	SecondaryFileInternetHost() *string
	SecondaryFileMicrosoftEndpoint() *string
	SecondaryFileMicrosoftHost() *string
	SecondaryLocation() *string
	SecondaryQueueEndpoint() *string
	SecondaryQueueHost() *string
	SecondaryQueueMicrosoftEndpoint() *string
	SecondaryQueueMicrosoftHost() *string
	SecondaryTableEndpoint() *string
	SecondaryTableHost() *string
	SecondaryTableMicrosoftEndpoint() *string
	SecondaryTableMicrosoftHost() *string
	SecondaryWebEndpoint() *string
	SecondaryWebHost() *string
	SecondaryWebInternetEndpoint() *string
	SecondaryWebInternetHost() *string
	SecondaryWebMicrosoftEndpoint() *string
	SecondaryWebMicrosoftHost() *string
	TableEncryptionKeyType() *string
	Tags() cdktf.StringMap
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAzurermStorageAccountTimeoutsOutputReference
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
	PutTimeouts(value *DataAzurermStorageAccountTimeouts)
	ResetId()
	ResetMinTlsVersion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAzurermStorageAccount
type jsiiProxy_DataAzurermStorageAccount struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzurermStorageAccount) AccessTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) AccountKind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) AccountReplicationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountReplicationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) AccountTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) AllowNestedItemsToBePublic() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowNestedItemsToBePublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) AzureFilesAuthentication() DataAzurermStorageAccountAzureFilesAuthenticationList {
	var returns DataAzurermStorageAccountAzureFilesAuthenticationList
	_jsii_.Get(
		j,
		"azureFilesAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) CustomDomain() DataAzurermStorageAccountCustomDomainList {
	var returns DataAzurermStorageAccountCustomDomainList
	_jsii_.Get(
		j,
		"customDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) DnsEndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsEndpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) HttpsTrafficOnlyEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"httpsTrafficOnlyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) Identity() DataAzurermStorageAccountIdentityList {
	var returns DataAzurermStorageAccountIdentityList
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) InfrastructureEncryptionEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"infrastructureEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) IsHnsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isHnsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) MinTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) MinTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) Nfsv3Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"nfsv3Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryBlobConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBlobConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryBlobEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBlobEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryBlobHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBlobHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryBlobInternetEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBlobInternetEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryBlobInternetHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBlobInternetHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryBlobMicrosoftEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBlobMicrosoftEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryBlobMicrosoftHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBlobMicrosoftHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryDfsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryDfsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryDfsHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryDfsHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryDfsInternetEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryDfsInternetEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryDfsInternetHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryDfsInternetHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryDfsMicrosoftEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryDfsMicrosoftEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryDfsMicrosoftHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryDfsMicrosoftHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryFileEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryFileEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryFileHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryFileHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryFileInternetEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryFileInternetEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryFileInternetHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryFileInternetHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryFileMicrosoftEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryFileMicrosoftEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryFileMicrosoftHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryFileMicrosoftHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryQueueEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryQueueEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryQueueHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryQueueHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryQueueMicrosoftEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryQueueMicrosoftEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryQueueMicrosoftHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryQueueMicrosoftHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryTableEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTableEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryTableHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTableHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryTableMicrosoftEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTableMicrosoftEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryTableMicrosoftHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTableMicrosoftHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryWebEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryWebEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryWebHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryWebHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryWebInternetEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryWebInternetEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryWebInternetHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryWebInternetHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryWebMicrosoftEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryWebMicrosoftEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) PrimaryWebMicrosoftHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryWebMicrosoftHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) QueueEncryptionKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueEncryptionKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryBlobConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBlobConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryBlobEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBlobEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryBlobHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBlobHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryBlobInternetEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBlobInternetEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryBlobInternetHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBlobInternetHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryBlobMicrosoftEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBlobMicrosoftEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryBlobMicrosoftHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBlobMicrosoftHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryDfsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryDfsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryDfsHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryDfsHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryDfsInternetEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryDfsInternetEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryDfsInternetHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryDfsInternetHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryDfsMicrosoftEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryDfsMicrosoftEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryDfsMicrosoftHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryDfsMicrosoftHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryFileEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryFileEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryFileHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryFileHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryFileInternetEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryFileInternetEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryFileInternetHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryFileInternetHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryFileMicrosoftEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryFileMicrosoftEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryFileMicrosoftHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryFileMicrosoftHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryQueueEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryQueueEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryQueueHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryQueueHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryQueueMicrosoftEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryQueueMicrosoftEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryQueueMicrosoftHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryQueueMicrosoftHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryTableEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryTableEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryTableHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryTableHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryTableMicrosoftEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryTableMicrosoftEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryTableMicrosoftHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryTableMicrosoftHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryWebEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryWebEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryWebHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryWebHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryWebInternetEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryWebInternetEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryWebInternetHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryWebInternetHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryWebMicrosoftEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryWebMicrosoftEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) SecondaryWebMicrosoftHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryWebMicrosoftHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) TableEncryptionKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableEncryptionKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) Tags() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) Timeouts() DataAzurermStorageAccountTimeoutsOutputReference {
	var returns DataAzurermStorageAccountTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccount) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/data-sources/storage_account azurerm_storage_account} Data Source.
func NewDataAzurermStorageAccount(scope constructs.Construct, id *string, config *DataAzurermStorageAccountConfig) DataAzurermStorageAccount {
	_init_.Initialize()

	if err := validateNewDataAzurermStorageAccountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermStorageAccount{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/data-sources/storage_account azurerm_storage_account} Data Source.
func NewDataAzurermStorageAccount_Override(d DataAzurermStorageAccount, scope constructs.Construct, id *string, config *DataAzurermStorageAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccount",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccount)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccount)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccount)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccount)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccount)SetMinTlsVersion(val *string) {
	if err := j.validateSetMinTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTlsVersion",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccount)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccount)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccount)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

// Generates CDKTF code for importing a DataAzurermStorageAccount resource upon running "cdktf plan <stack-name>".
func DataAzurermStorageAccount_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAzurermStorageAccount_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccount",
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
func DataAzurermStorageAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermStorageAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermStorageAccount_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermStorageAccount_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccount",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermStorageAccount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermStorageAccount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzurermStorageAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataAzurermStorageAccount.DataAzurermStorageAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermStorageAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermStorageAccount) PutTimeouts(value *DataAzurermStorageAccountTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermStorageAccount) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermStorageAccount) ResetMinTlsVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetMinTlsVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermStorageAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermStorageAccount) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermStorageAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

