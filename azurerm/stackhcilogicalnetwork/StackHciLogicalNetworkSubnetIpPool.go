// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcilogicalnetwork


type StackHciLogicalNetworkSubnetIpPool struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/stack_hci_logical_network#end StackHciLogicalNetwork#end}.
	End *string `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/stack_hci_logical_network#start StackHciLogicalNetwork#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
}

