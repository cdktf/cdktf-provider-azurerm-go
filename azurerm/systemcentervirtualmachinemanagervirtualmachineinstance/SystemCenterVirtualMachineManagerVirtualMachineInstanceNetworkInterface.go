// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package systemcentervirtualmachinemanagervirtualmachineinstance


type SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#name SystemCenterVirtualMachineManagerVirtualMachineInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#ipv4_address_type SystemCenterVirtualMachineManagerVirtualMachineInstance#ipv4_address_type}.
	Ipv4AddressType *string `field:"optional" json:"ipv4AddressType" yaml:"ipv4AddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#ipv6_address_type SystemCenterVirtualMachineManagerVirtualMachineInstance#ipv6_address_type}.
	Ipv6AddressType *string `field:"optional" json:"ipv6AddressType" yaml:"ipv6AddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#mac_address_type SystemCenterVirtualMachineManagerVirtualMachineInstance#mac_address_type}.
	MacAddressType *string `field:"optional" json:"macAddressType" yaml:"macAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance#virtual_network_id SystemCenterVirtualMachineManagerVirtualMachineInstance#virtual_network_id}.
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

