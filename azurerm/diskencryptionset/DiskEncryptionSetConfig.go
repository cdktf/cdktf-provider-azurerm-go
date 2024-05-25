// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package diskencryptionset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DiskEncryptionSetConfig struct {
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
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/disk_encryption_set#identity DiskEncryptionSet#identity}
	Identity *DiskEncryptionSetIdentity `field:"required" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/disk_encryption_set#key_vault_key_id DiskEncryptionSet#key_vault_key_id}.
	KeyVaultKeyId *string `field:"required" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/disk_encryption_set#location DiskEncryptionSet#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/disk_encryption_set#name DiskEncryptionSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/disk_encryption_set#resource_group_name DiskEncryptionSet#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/disk_encryption_set#auto_key_rotation_enabled DiskEncryptionSet#auto_key_rotation_enabled}.
	AutoKeyRotationEnabled interface{} `field:"optional" json:"autoKeyRotationEnabled" yaml:"autoKeyRotationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/disk_encryption_set#encryption_type DiskEncryptionSet#encryption_type}.
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/disk_encryption_set#federated_client_id DiskEncryptionSet#federated_client_id}.
	FederatedClientId *string `field:"optional" json:"federatedClientId" yaml:"federatedClientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/disk_encryption_set#id DiskEncryptionSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/disk_encryption_set#tags DiskEncryptionSet#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/disk_encryption_set#timeouts DiskEncryptionSet#timeouts}
	Timeouts *DiskEncryptionSetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

