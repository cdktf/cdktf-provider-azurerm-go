// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwarenetappvolumeattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmwareNetappVolumeAttachmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/vmware_netapp_volume_attachment#name VmwareNetappVolumeAttachment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/vmware_netapp_volume_attachment#netapp_volume_id VmwareNetappVolumeAttachment#netapp_volume_id}.
	NetappVolumeId *string `field:"required" json:"netappVolumeId" yaml:"netappVolumeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/vmware_netapp_volume_attachment#vmware_cluster_id VmwareNetappVolumeAttachment#vmware_cluster_id}.
	VmwareClusterId *string `field:"required" json:"vmwareClusterId" yaml:"vmwareClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/vmware_netapp_volume_attachment#id VmwareNetappVolumeAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/vmware_netapp_volume_attachment#timeouts VmwareNetappVolumeAttachment#timeouts}
	Timeouts *VmwareNetappVolumeAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

