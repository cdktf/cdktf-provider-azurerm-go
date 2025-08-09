// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermsystemcentervirtualmachinemanagerinventoryitems

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermSystemCenterVirtualMachineManagerInventoryItemsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/data-sources/system_center_virtual_machine_manager_inventory_items#inventory_type DataAzurermSystemCenterVirtualMachineManagerInventoryItems#inventory_type}.
	InventoryType *string `field:"required" json:"inventoryType" yaml:"inventoryType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/data-sources/system_center_virtual_machine_manager_inventory_items#system_center_virtual_machine_manager_server_id DataAzurermSystemCenterVirtualMachineManagerInventoryItems#system_center_virtual_machine_manager_server_id}.
	SystemCenterVirtualMachineManagerServerId *string `field:"required" json:"systemCenterVirtualMachineManagerServerId" yaml:"systemCenterVirtualMachineManagerServerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/data-sources/system_center_virtual_machine_manager_inventory_items#id DataAzurermSystemCenterVirtualMachineManagerInventoryItems#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/data-sources/system_center_virtual_machine_manager_inventory_items#timeouts DataAzurermSystemCenterVirtualMachineManagerInventoryItems#timeouts}
	Timeouts *DataAzurermSystemCenterVirtualMachineManagerInventoryItemsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

