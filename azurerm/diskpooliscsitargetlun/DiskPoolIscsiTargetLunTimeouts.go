// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package diskpooliscsitargetlun


type DiskPoolIscsiTargetLunTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/disk_pool_iscsi_target_lun#create DiskPoolIscsiTargetLun#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/disk_pool_iscsi_target_lun#delete DiskPoolIscsiTargetLun#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/disk_pool_iscsi_target_lun#read DiskPoolIscsiTargetLun#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

