// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcilogicalnetwork


type StackHciLogicalNetworkTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/stack_hci_logical_network#create StackHciLogicalNetwork#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/stack_hci_logical_network#delete StackHciLogicalNetwork#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/stack_hci_logical_network#read StackHciLogicalNetwork#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/stack_hci_logical_network#update StackHciLogicalNetwork#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

