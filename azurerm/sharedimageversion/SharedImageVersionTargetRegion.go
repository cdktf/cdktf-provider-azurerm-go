// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sharedimageversion


type SharedImageVersionTargetRegion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/shared_image_version#name SharedImageVersion#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/shared_image_version#regional_replica_count SharedImageVersion#regional_replica_count}.
	RegionalReplicaCount *float64 `field:"required" json:"regionalReplicaCount" yaml:"regionalReplicaCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/shared_image_version#disk_encryption_set_id SharedImageVersion#disk_encryption_set_id}.
	DiskEncryptionSetId *string `field:"optional" json:"diskEncryptionSetId" yaml:"diskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/shared_image_version#storage_account_type SharedImageVersion#storage_account_type}.
	StorageAccountType *string `field:"optional" json:"storageAccountType" yaml:"storageAccountType"`
}

