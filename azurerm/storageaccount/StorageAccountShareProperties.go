// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount


type StorageAccountShareProperties struct {
	// cors_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/storage_account#cors_rule StorageAccount#cors_rule}
	CorsRule interface{} `field:"optional" json:"corsRule" yaml:"corsRule"`
	// retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/storage_account#retention_policy StorageAccount#retention_policy}
	RetentionPolicy *StorageAccountSharePropertiesRetentionPolicy `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
	// smb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/storage_account#smb StorageAccount#smb}
	Smb *StorageAccountSharePropertiesSmb `field:"optional" json:"smb" yaml:"smb"`
}

