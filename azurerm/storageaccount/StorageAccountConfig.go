// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageAccountConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#account_replication_type StorageAccount#account_replication_type}.
	AccountReplicationType *string `field:"required" json:"accountReplicationType" yaml:"accountReplicationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#account_tier StorageAccount#account_tier}.
	AccountTier *string `field:"required" json:"accountTier" yaml:"accountTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#location StorageAccount#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#name StorageAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#resource_group_name StorageAccount#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#access_tier StorageAccount#access_tier}.
	AccessTier *string `field:"optional" json:"accessTier" yaml:"accessTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#account_kind StorageAccount#account_kind}.
	AccountKind *string `field:"optional" json:"accountKind" yaml:"accountKind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#allowed_copy_scope StorageAccount#allowed_copy_scope}.
	AllowedCopyScope *string `field:"optional" json:"allowedCopyScope" yaml:"allowedCopyScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#allow_nested_items_to_be_public StorageAccount#allow_nested_items_to_be_public}.
	AllowNestedItemsToBePublic interface{} `field:"optional" json:"allowNestedItemsToBePublic" yaml:"allowNestedItemsToBePublic"`
	// azure_files_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#azure_files_authentication StorageAccount#azure_files_authentication}
	AzureFilesAuthentication *StorageAccountAzureFilesAuthentication `field:"optional" json:"azureFilesAuthentication" yaml:"azureFilesAuthentication"`
	// blob_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#blob_properties StorageAccount#blob_properties}
	BlobProperties *StorageAccountBlobProperties `field:"optional" json:"blobProperties" yaml:"blobProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#cross_tenant_replication_enabled StorageAccount#cross_tenant_replication_enabled}.
	CrossTenantReplicationEnabled interface{} `field:"optional" json:"crossTenantReplicationEnabled" yaml:"crossTenantReplicationEnabled"`
	// custom_domain block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#custom_domain StorageAccount#custom_domain}
	CustomDomain *StorageAccountCustomDomain `field:"optional" json:"customDomain" yaml:"customDomain"`
	// customer_managed_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#customer_managed_key StorageAccount#customer_managed_key}
	CustomerManagedKey *StorageAccountCustomerManagedKey `field:"optional" json:"customerManagedKey" yaml:"customerManagedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#default_to_oauth_authentication StorageAccount#default_to_oauth_authentication}.
	DefaultToOauthAuthentication interface{} `field:"optional" json:"defaultToOauthAuthentication" yaml:"defaultToOauthAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#dns_endpoint_type StorageAccount#dns_endpoint_type}.
	DnsEndpointType *string `field:"optional" json:"dnsEndpointType" yaml:"dnsEndpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#edge_zone StorageAccount#edge_zone}.
	EdgeZone *string `field:"optional" json:"edgeZone" yaml:"edgeZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#https_traffic_only_enabled StorageAccount#https_traffic_only_enabled}.
	HttpsTrafficOnlyEnabled interface{} `field:"optional" json:"httpsTrafficOnlyEnabled" yaml:"httpsTrafficOnlyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#id StorageAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#identity StorageAccount#identity}
	Identity *StorageAccountIdentity `field:"optional" json:"identity" yaml:"identity"`
	// immutability_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#immutability_policy StorageAccount#immutability_policy}
	ImmutabilityPolicy *StorageAccountImmutabilityPolicy `field:"optional" json:"immutabilityPolicy" yaml:"immutabilityPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#infrastructure_encryption_enabled StorageAccount#infrastructure_encryption_enabled}.
	InfrastructureEncryptionEnabled interface{} `field:"optional" json:"infrastructureEncryptionEnabled" yaml:"infrastructureEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#is_hns_enabled StorageAccount#is_hns_enabled}.
	IsHnsEnabled interface{} `field:"optional" json:"isHnsEnabled" yaml:"isHnsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#large_file_share_enabled StorageAccount#large_file_share_enabled}.
	LargeFileShareEnabled interface{} `field:"optional" json:"largeFileShareEnabled" yaml:"largeFileShareEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#local_user_enabled StorageAccount#local_user_enabled}.
	LocalUserEnabled interface{} `field:"optional" json:"localUserEnabled" yaml:"localUserEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#min_tls_version StorageAccount#min_tls_version}.
	MinTlsVersion *string `field:"optional" json:"minTlsVersion" yaml:"minTlsVersion"`
	// network_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#network_rules StorageAccount#network_rules}
	NetworkRules *StorageAccountNetworkRules `field:"optional" json:"networkRules" yaml:"networkRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#nfsv3_enabled StorageAccount#nfsv3_enabled}.
	Nfsv3Enabled interface{} `field:"optional" json:"nfsv3Enabled" yaml:"nfsv3Enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#public_network_access_enabled StorageAccount#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#queue_encryption_key_type StorageAccount#queue_encryption_key_type}.
	QueueEncryptionKeyType *string `field:"optional" json:"queueEncryptionKeyType" yaml:"queueEncryptionKeyType"`
	// queue_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#queue_properties StorageAccount#queue_properties}
	QueueProperties *StorageAccountQueueProperties `field:"optional" json:"queueProperties" yaml:"queueProperties"`
	// routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#routing StorageAccount#routing}
	Routing *StorageAccountRouting `field:"optional" json:"routing" yaml:"routing"`
	// sas_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#sas_policy StorageAccount#sas_policy}
	SasPolicy *StorageAccountSasPolicy `field:"optional" json:"sasPolicy" yaml:"sasPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#sftp_enabled StorageAccount#sftp_enabled}.
	SftpEnabled interface{} `field:"optional" json:"sftpEnabled" yaml:"sftpEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#shared_access_key_enabled StorageAccount#shared_access_key_enabled}.
	SharedAccessKeyEnabled interface{} `field:"optional" json:"sharedAccessKeyEnabled" yaml:"sharedAccessKeyEnabled"`
	// share_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#share_properties StorageAccount#share_properties}
	ShareProperties *StorageAccountShareProperties `field:"optional" json:"shareProperties" yaml:"shareProperties"`
	// static_website block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#static_website StorageAccount#static_website}
	StaticWebsite *StorageAccountStaticWebsite `field:"optional" json:"staticWebsite" yaml:"staticWebsite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#table_encryption_key_type StorageAccount#table_encryption_key_type}.
	TableEncryptionKeyType *string `field:"optional" json:"tableEncryptionKeyType" yaml:"tableEncryptionKeyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#tags StorageAccount#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/storage_account#timeouts StorageAccount#timeouts}
	Timeouts *StorageAccountTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

