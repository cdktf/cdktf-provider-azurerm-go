// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescaleset


type VirtualMachineScaleSetStorageProfileOsDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_machine_scale_set#create_option VirtualMachineScaleSet#create_option}.
	CreateOption *string `field:"required" json:"createOption" yaml:"createOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_machine_scale_set#caching VirtualMachineScaleSet#caching}.
	Caching *string `field:"optional" json:"caching" yaml:"caching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_machine_scale_set#image VirtualMachineScaleSet#image}.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_machine_scale_set#managed_disk_type VirtualMachineScaleSet#managed_disk_type}.
	ManagedDiskType *string `field:"optional" json:"managedDiskType" yaml:"managedDiskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_machine_scale_set#name VirtualMachineScaleSet#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_machine_scale_set#os_type VirtualMachineScaleSet#os_type}.
	OsType *string `field:"optional" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_machine_scale_set#vhd_containers VirtualMachineScaleSet#vhd_containers}.
	VhdContainers *[]*string `field:"optional" json:"vhdContainers" yaml:"vhdContainers"`
}

