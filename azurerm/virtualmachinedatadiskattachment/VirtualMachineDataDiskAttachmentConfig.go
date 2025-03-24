// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinedatadiskattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineDataDiskAttachmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_machine_data_disk_attachment#caching VirtualMachineDataDiskAttachment#caching}.
	Caching *string `field:"required" json:"caching" yaml:"caching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_machine_data_disk_attachment#lun VirtualMachineDataDiskAttachment#lun}.
	Lun *float64 `field:"required" json:"lun" yaml:"lun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_machine_data_disk_attachment#managed_disk_id VirtualMachineDataDiskAttachment#managed_disk_id}.
	ManagedDiskId *string `field:"required" json:"managedDiskId" yaml:"managedDiskId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_machine_data_disk_attachment#virtual_machine_id VirtualMachineDataDiskAttachment#virtual_machine_id}.
	VirtualMachineId *string `field:"required" json:"virtualMachineId" yaml:"virtualMachineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_machine_data_disk_attachment#create_option VirtualMachineDataDiskAttachment#create_option}.
	CreateOption *string `field:"optional" json:"createOption" yaml:"createOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_machine_data_disk_attachment#id VirtualMachineDataDiskAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_machine_data_disk_attachment#timeouts VirtualMachineDataDiskAttachment#timeouts}
	Timeouts *VirtualMachineDataDiskAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_machine_data_disk_attachment#write_accelerator_enabled VirtualMachineDataDiskAttachment#write_accelerator_enabled}.
	WriteAcceleratorEnabled interface{} `field:"optional" json:"writeAcceleratorEnabled" yaml:"writeAcceleratorEnabled"`
}

