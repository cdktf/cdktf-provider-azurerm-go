// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package systemcentervirtualmachinemanagervirtualnetwork

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SystemCenterVirtualMachineManagerVirtualNetworkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/system_center_virtual_machine_manager_virtual_network#custom_location_id SystemCenterVirtualMachineManagerVirtualNetwork#custom_location_id}.
	CustomLocationId *string `field:"required" json:"customLocationId" yaml:"customLocationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/system_center_virtual_machine_manager_virtual_network#location SystemCenterVirtualMachineManagerVirtualNetwork#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/system_center_virtual_machine_manager_virtual_network#name SystemCenterVirtualMachineManagerVirtualNetwork#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/system_center_virtual_machine_manager_virtual_network#resource_group_name SystemCenterVirtualMachineManagerVirtualNetwork#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/system_center_virtual_machine_manager_virtual_network#system_center_virtual_machine_manager_server_inventory_item_id SystemCenterVirtualMachineManagerVirtualNetwork#system_center_virtual_machine_manager_server_inventory_item_id}.
	SystemCenterVirtualMachineManagerServerInventoryItemId *string `field:"required" json:"systemCenterVirtualMachineManagerServerInventoryItemId" yaml:"systemCenterVirtualMachineManagerServerInventoryItemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/system_center_virtual_machine_manager_virtual_network#id SystemCenterVirtualMachineManagerVirtualNetwork#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/system_center_virtual_machine_manager_virtual_network#tags SystemCenterVirtualMachineManagerVirtualNetwork#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/system_center_virtual_machine_manager_virtual_network#timeouts SystemCenterVirtualMachineManagerVirtualNetwork#timeouts}
	Timeouts *SystemCenterVirtualMachineManagerVirtualNetworkTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

