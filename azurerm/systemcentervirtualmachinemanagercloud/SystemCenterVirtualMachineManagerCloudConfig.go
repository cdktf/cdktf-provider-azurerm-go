// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package systemcentervirtualmachinemanagercloud

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SystemCenterVirtualMachineManagerCloudConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/system_center_virtual_machine_manager_cloud#custom_location_id SystemCenterVirtualMachineManagerCloud#custom_location_id}.
	CustomLocationId *string `field:"required" json:"customLocationId" yaml:"customLocationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/system_center_virtual_machine_manager_cloud#location SystemCenterVirtualMachineManagerCloud#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/system_center_virtual_machine_manager_cloud#name SystemCenterVirtualMachineManagerCloud#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/system_center_virtual_machine_manager_cloud#resource_group_name SystemCenterVirtualMachineManagerCloud#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/system_center_virtual_machine_manager_cloud#system_center_virtual_machine_manager_server_inventory_item_id SystemCenterVirtualMachineManagerCloud#system_center_virtual_machine_manager_server_inventory_item_id}.
	SystemCenterVirtualMachineManagerServerInventoryItemId *string `field:"required" json:"systemCenterVirtualMachineManagerServerInventoryItemId" yaml:"systemCenterVirtualMachineManagerServerInventoryItemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/system_center_virtual_machine_manager_cloud#id SystemCenterVirtualMachineManagerCloud#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/system_center_virtual_machine_manager_cloud#tags SystemCenterVirtualMachineManagerCloud#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/system_center_virtual_machine_manager_cloud#timeouts SystemCenterVirtualMachineManagerCloud#timeouts}
	Timeouts *SystemCenterVirtualMachineManagerCloudTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

