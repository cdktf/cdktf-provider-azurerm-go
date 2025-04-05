// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package systemcentervirtualmachinemanagervirtualmachineinstance


type SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#bus SystemCenterVirtualMachineManagerVirtualMachineInstance#bus}.
	Bus *float64 `field:"optional" json:"bus" yaml:"bus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#bus_type SystemCenterVirtualMachineManagerVirtualMachineInstance#bus_type}.
	BusType *string `field:"optional" json:"busType" yaml:"busType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#disk_size_gb SystemCenterVirtualMachineManagerVirtualMachineInstance#disk_size_gb}.
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#lun SystemCenterVirtualMachineManagerVirtualMachineInstance#lun}.
	Lun *float64 `field:"optional" json:"lun" yaml:"lun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#name SystemCenterVirtualMachineManagerVirtualMachineInstance#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#storage_qos_policy_name SystemCenterVirtualMachineManagerVirtualMachineInstance#storage_qos_policy_name}.
	StorageQosPolicyName *string `field:"optional" json:"storageQosPolicyName" yaml:"storageQosPolicyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#template_disk_id SystemCenterVirtualMachineManagerVirtualMachineInstance#template_disk_id}.
	TemplateDiskId *string `field:"optional" json:"templateDiskId" yaml:"templateDiskId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#vhd_type SystemCenterVirtualMachineManagerVirtualMachineInstance#vhd_type}.
	VhdType *string `field:"optional" json:"vhdType" yaml:"vhdType"`
}

