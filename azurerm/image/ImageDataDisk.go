// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package image


type ImageDataDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/image#blob_uri Image#blob_uri}.
	BlobUri *string `field:"optional" json:"blobUri" yaml:"blobUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/image#caching Image#caching}.
	Caching *string `field:"optional" json:"caching" yaml:"caching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/image#lun Image#lun}.
	Lun *float64 `field:"optional" json:"lun" yaml:"lun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/image#managed_disk_id Image#managed_disk_id}.
	ManagedDiskId *string `field:"optional" json:"managedDiskId" yaml:"managedDiskId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/image#size_gb Image#size_gb}.
	SizeGb *float64 `field:"optional" json:"sizeGb" yaml:"sizeGb"`
}

