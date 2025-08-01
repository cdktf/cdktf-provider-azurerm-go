// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package image


type ImageOsDisk struct {
	// The type of storage disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/image#storage_type Image#storage_type}
	StorageType *string `field:"required" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/image#blob_uri Image#blob_uri}.
	BlobUri *string `field:"optional" json:"blobUri" yaml:"blobUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/image#caching Image#caching}.
	Caching *string `field:"optional" json:"caching" yaml:"caching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/image#disk_encryption_set_id Image#disk_encryption_set_id}.
	DiskEncryptionSetId *string `field:"optional" json:"diskEncryptionSetId" yaml:"diskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/image#managed_disk_id Image#managed_disk_id}.
	ManagedDiskId *string `field:"optional" json:"managedDiskId" yaml:"managedDiskId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/image#os_state Image#os_state}.
	OsState *string `field:"optional" json:"osState" yaml:"osState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/image#os_type Image#os_type}.
	OsType *string `field:"optional" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/image#size_gb Image#size_gb}.
	SizeGb *float64 `field:"optional" json:"sizeGb" yaml:"sizeGb"`
}

