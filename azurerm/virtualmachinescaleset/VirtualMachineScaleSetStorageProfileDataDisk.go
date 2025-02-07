// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescaleset


type VirtualMachineScaleSetStorageProfileDataDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/virtual_machine_scale_set#create_option VirtualMachineScaleSet#create_option}.
	CreateOption *string `field:"required" json:"createOption" yaml:"createOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/virtual_machine_scale_set#lun VirtualMachineScaleSet#lun}.
	Lun *float64 `field:"required" json:"lun" yaml:"lun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/virtual_machine_scale_set#caching VirtualMachineScaleSet#caching}.
	Caching *string `field:"optional" json:"caching" yaml:"caching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/virtual_machine_scale_set#disk_size_gb VirtualMachineScaleSet#disk_size_gb}.
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/virtual_machine_scale_set#managed_disk_type VirtualMachineScaleSet#managed_disk_type}.
	ManagedDiskType *string `field:"optional" json:"managedDiskType" yaml:"managedDiskType"`
}

