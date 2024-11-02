// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagesynccloudendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageSyncCloudEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/storage_sync_cloud_endpoint#file_share_name StorageSyncCloudEndpoint#file_share_name}.
	FileShareName *string `field:"required" json:"fileShareName" yaml:"fileShareName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/storage_sync_cloud_endpoint#name StorageSyncCloudEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/storage_sync_cloud_endpoint#storage_account_id StorageSyncCloudEndpoint#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/storage_sync_cloud_endpoint#storage_sync_group_id StorageSyncCloudEndpoint#storage_sync_group_id}.
	StorageSyncGroupId *string `field:"required" json:"storageSyncGroupId" yaml:"storageSyncGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/storage_sync_cloud_endpoint#id StorageSyncCloudEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/storage_sync_cloud_endpoint#storage_account_tenant_id StorageSyncCloudEndpoint#storage_account_tenant_id}.
	StorageAccountTenantId *string `field:"optional" json:"storageAccountTenantId" yaml:"storageAccountTenantId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/storage_sync_cloud_endpoint#timeouts StorageSyncCloudEndpoint#timeouts}
	Timeouts *StorageSyncCloudEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

