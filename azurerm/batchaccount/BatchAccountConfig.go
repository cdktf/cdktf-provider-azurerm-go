// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchAccountConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#location BatchAccount#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#name BatchAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#resource_group_name BatchAccount#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#allowed_authentication_modes BatchAccount#allowed_authentication_modes}.
	AllowedAuthenticationModes *[]*string `field:"optional" json:"allowedAuthenticationModes" yaml:"allowedAuthenticationModes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#encryption BatchAccount#encryption}.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#id BatchAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#identity BatchAccount#identity}
	Identity *BatchAccountIdentity `field:"optional" json:"identity" yaml:"identity"`
	// key_vault_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#key_vault_reference BatchAccount#key_vault_reference}
	KeyVaultReference *BatchAccountKeyVaultReference `field:"optional" json:"keyVaultReference" yaml:"keyVaultReference"`
	// network_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#network_profile BatchAccount#network_profile}
	NetworkProfile *BatchAccountNetworkProfile `field:"optional" json:"networkProfile" yaml:"networkProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#pool_allocation_mode BatchAccount#pool_allocation_mode}.
	PoolAllocationMode *string `field:"optional" json:"poolAllocationMode" yaml:"poolAllocationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#public_network_access_enabled BatchAccount#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#storage_account_authentication_mode BatchAccount#storage_account_authentication_mode}.
	StorageAccountAuthenticationMode *string `field:"optional" json:"storageAccountAuthenticationMode" yaml:"storageAccountAuthenticationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#storage_account_id BatchAccount#storage_account_id}.
	StorageAccountId *string `field:"optional" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#storage_account_node_identity BatchAccount#storage_account_node_identity}.
	StorageAccountNodeIdentity *string `field:"optional" json:"storageAccountNodeIdentity" yaml:"storageAccountNodeIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#tags BatchAccount#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/batch_account#timeouts BatchAccount#timeouts}
	Timeouts *BatchAccountTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

