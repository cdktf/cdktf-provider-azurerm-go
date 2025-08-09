// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package systemcentervirtualmachinemanagervirtualmachineinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SystemCenterVirtualMachineManagerVirtualMachineInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#custom_location_id SystemCenterVirtualMachineManagerVirtualMachineInstance#custom_location_id}.
	CustomLocationId *string `field:"required" json:"customLocationId" yaml:"customLocationId"`
	// infrastructure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#infrastructure SystemCenterVirtualMachineManagerVirtualMachineInstance#infrastructure}
	Infrastructure *SystemCenterVirtualMachineManagerVirtualMachineInstanceInfrastructure `field:"required" json:"infrastructure" yaml:"infrastructure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#scoped_resource_id SystemCenterVirtualMachineManagerVirtualMachineInstance#scoped_resource_id}.
	ScopedResourceId *string `field:"required" json:"scopedResourceId" yaml:"scopedResourceId"`
	// hardware block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#hardware SystemCenterVirtualMachineManagerVirtualMachineInstance#hardware}
	Hardware *SystemCenterVirtualMachineManagerVirtualMachineInstanceHardware `field:"optional" json:"hardware" yaml:"hardware"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#id SystemCenterVirtualMachineManagerVirtualMachineInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#network_interface SystemCenterVirtualMachineManagerVirtualMachineInstance#network_interface}
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// operating_system block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#operating_system SystemCenterVirtualMachineManagerVirtualMachineInstance#operating_system}
	OperatingSystem *SystemCenterVirtualMachineManagerVirtualMachineInstanceOperatingSystem `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// storage_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#storage_disk SystemCenterVirtualMachineManagerVirtualMachineInstance#storage_disk}
	StorageDisk interface{} `field:"optional" json:"storageDisk" yaml:"storageDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#system_center_virtual_machine_manager_availability_set_ids SystemCenterVirtualMachineManagerVirtualMachineInstance#system_center_virtual_machine_manager_availability_set_ids}.
	SystemCenterVirtualMachineManagerAvailabilitySetIds *[]*string `field:"optional" json:"systemCenterVirtualMachineManagerAvailabilitySetIds" yaml:"systemCenterVirtualMachineManagerAvailabilitySetIds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#timeouts SystemCenterVirtualMachineManagerVirtualMachineInstance#timeouts}
	Timeouts *SystemCenterVirtualMachineManagerVirtualMachineInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

