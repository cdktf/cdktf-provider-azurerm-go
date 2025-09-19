// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcilogicalnetwork


type StackHciLogicalNetworkSubnet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/stack_hci_logical_network#ip_allocation_method StackHciLogicalNetwork#ip_allocation_method}.
	IpAllocationMethod *string `field:"required" json:"ipAllocationMethod" yaml:"ipAllocationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/stack_hci_logical_network#address_prefix StackHciLogicalNetwork#address_prefix}.
	AddressPrefix *string `field:"optional" json:"addressPrefix" yaml:"addressPrefix"`
	// ip_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/stack_hci_logical_network#ip_pool StackHciLogicalNetwork#ip_pool}
	IpPool interface{} `field:"optional" json:"ipPool" yaml:"ipPool"`
	// route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/stack_hci_logical_network#route StackHciLogicalNetwork#route}
	Route *StackHciLogicalNetworkSubnetRoute `field:"optional" json:"route" yaml:"route"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/stack_hci_logical_network#vlan_id StackHciLogicalNetwork#vlan_id}.
	VlanId *float64 `field:"optional" json:"vlanId" yaml:"vlanId"`
}

