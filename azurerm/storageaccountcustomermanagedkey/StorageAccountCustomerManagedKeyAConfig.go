// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccountcustomermanagedkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageAccountCustomerManagedKeyAConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/storage_account_customer_managed_key#key_name StorageAccountCustomerManagedKeyA#key_name}.
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/storage_account_customer_managed_key#storage_account_id StorageAccountCustomerManagedKeyA#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/storage_account_customer_managed_key#federated_identity_client_id StorageAccountCustomerManagedKeyA#federated_identity_client_id}.
	FederatedIdentityClientId *string `field:"optional" json:"federatedIdentityClientId" yaml:"federatedIdentityClientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/storage_account_customer_managed_key#id StorageAccountCustomerManagedKeyA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/storage_account_customer_managed_key#key_vault_id StorageAccountCustomerManagedKeyA#key_vault_id}.
	KeyVaultId *string `field:"optional" json:"keyVaultId" yaml:"keyVaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/storage_account_customer_managed_key#key_vault_uri StorageAccountCustomerManagedKeyA#key_vault_uri}.
	KeyVaultUri *string `field:"optional" json:"keyVaultUri" yaml:"keyVaultUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/storage_account_customer_managed_key#key_version StorageAccountCustomerManagedKeyA#key_version}.
	KeyVersion *string `field:"optional" json:"keyVersion" yaml:"keyVersion"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/storage_account_customer_managed_key#timeouts StorageAccountCustomerManagedKeyA#timeouts}
	Timeouts *StorageAccountCustomerManagedKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/storage_account_customer_managed_key#user_assigned_identity_id StorageAccountCustomerManagedKeyA#user_assigned_identity_id}.
	UserAssignedIdentityId *string `field:"optional" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
}

