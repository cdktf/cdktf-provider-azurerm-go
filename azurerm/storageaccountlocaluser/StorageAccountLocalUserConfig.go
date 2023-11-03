// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccountlocaluser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageAccountLocalUserConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/storage_account_local_user#name StorageAccountLocalUser#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/storage_account_local_user#storage_account_id StorageAccountLocalUser#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/storage_account_local_user#home_directory StorageAccountLocalUser#home_directory}.
	HomeDirectory *string `field:"optional" json:"homeDirectory" yaml:"homeDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/storage_account_local_user#id StorageAccountLocalUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// permission_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/storage_account_local_user#permission_scope StorageAccountLocalUser#permission_scope}
	PermissionScope interface{} `field:"optional" json:"permissionScope" yaml:"permissionScope"`
	// ssh_authorized_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/storage_account_local_user#ssh_authorized_key StorageAccountLocalUser#ssh_authorized_key}
	SshAuthorizedKey interface{} `field:"optional" json:"sshAuthorizedKey" yaml:"sshAuthorizedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/storage_account_local_user#ssh_key_enabled StorageAccountLocalUser#ssh_key_enabled}.
	SshKeyEnabled interface{} `field:"optional" json:"sshKeyEnabled" yaml:"sshKeyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/storage_account_local_user#ssh_password_enabled StorageAccountLocalUser#ssh_password_enabled}.
	SshPasswordEnabled interface{} `field:"optional" json:"sshPasswordEnabled" yaml:"sshPasswordEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/storage_account_local_user#timeouts StorageAccountLocalUser#timeouts}
	Timeouts *StorageAccountLocalUserTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

