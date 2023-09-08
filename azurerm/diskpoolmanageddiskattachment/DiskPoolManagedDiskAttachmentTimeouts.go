// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package diskpoolmanageddiskattachment


type DiskPoolManagedDiskAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/disk_pool_managed_disk_attachment#create DiskPoolManagedDiskAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/disk_pool_managed_disk_attachment#delete DiskPoolManagedDiskAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/disk_pool_managed_disk_attachment#read DiskPoolManagedDiskAttachment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

