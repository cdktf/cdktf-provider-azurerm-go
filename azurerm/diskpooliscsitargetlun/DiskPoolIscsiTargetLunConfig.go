// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package diskpooliscsitargetlun

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DiskPoolIscsiTargetLunConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/disk_pool_iscsi_target_lun#disk_pool_managed_disk_attachment_id DiskPoolIscsiTargetLun#disk_pool_managed_disk_attachment_id}.
	DiskPoolManagedDiskAttachmentId *string `field:"required" json:"diskPoolManagedDiskAttachmentId" yaml:"diskPoolManagedDiskAttachmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/disk_pool_iscsi_target_lun#iscsi_target_id DiskPoolIscsiTargetLun#iscsi_target_id}.
	IscsiTargetId *string `field:"required" json:"iscsiTargetId" yaml:"iscsiTargetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/disk_pool_iscsi_target_lun#name DiskPoolIscsiTargetLun#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/disk_pool_iscsi_target_lun#id DiskPoolIscsiTargetLun#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/disk_pool_iscsi_target_lun#timeouts DiskPoolIscsiTargetLun#timeouts}
	Timeouts *DiskPoolIscsiTargetLunTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

