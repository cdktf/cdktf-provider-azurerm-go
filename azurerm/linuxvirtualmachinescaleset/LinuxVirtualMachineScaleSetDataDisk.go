// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetDataDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/linux_virtual_machine_scale_set#caching LinuxVirtualMachineScaleSet#caching}.
	Caching *string `field:"required" json:"caching" yaml:"caching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/linux_virtual_machine_scale_set#disk_size_gb LinuxVirtualMachineScaleSet#disk_size_gb}.
	DiskSizeGb *float64 `field:"required" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/linux_virtual_machine_scale_set#lun LinuxVirtualMachineScaleSet#lun}.
	Lun *float64 `field:"required" json:"lun" yaml:"lun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/linux_virtual_machine_scale_set#storage_account_type LinuxVirtualMachineScaleSet#storage_account_type}.
	StorageAccountType *string `field:"required" json:"storageAccountType" yaml:"storageAccountType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/linux_virtual_machine_scale_set#create_option LinuxVirtualMachineScaleSet#create_option}.
	CreateOption *string `field:"optional" json:"createOption" yaml:"createOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/linux_virtual_machine_scale_set#disk_encryption_set_id LinuxVirtualMachineScaleSet#disk_encryption_set_id}.
	DiskEncryptionSetId *string `field:"optional" json:"diskEncryptionSetId" yaml:"diskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/linux_virtual_machine_scale_set#name LinuxVirtualMachineScaleSet#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/linux_virtual_machine_scale_set#ultra_ssd_disk_iops_read_write LinuxVirtualMachineScaleSet#ultra_ssd_disk_iops_read_write}.
	UltraSsdDiskIopsReadWrite *float64 `field:"optional" json:"ultraSsdDiskIopsReadWrite" yaml:"ultraSsdDiskIopsReadWrite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/linux_virtual_machine_scale_set#ultra_ssd_disk_mbps_read_write LinuxVirtualMachineScaleSet#ultra_ssd_disk_mbps_read_write}.
	UltraSsdDiskMbpsReadWrite *float64 `field:"optional" json:"ultraSsdDiskMbpsReadWrite" yaml:"ultraSsdDiskMbpsReadWrite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/linux_virtual_machine_scale_set#write_accelerator_enabled LinuxVirtualMachineScaleSet#write_accelerator_enabled}.
	WriteAcceleratorEnabled interface{} `field:"optional" json:"writeAcceleratorEnabled" yaml:"writeAcceleratorEnabled"`
}

