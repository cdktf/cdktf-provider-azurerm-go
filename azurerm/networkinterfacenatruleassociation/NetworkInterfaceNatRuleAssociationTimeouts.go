// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkinterfacenatruleassociation


type NetworkInterfaceNatRuleAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/network_interface_nat_rule_association#create NetworkInterfaceNatRuleAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/network_interface_nat_rule_association#delete NetworkInterfaceNatRuleAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/network_interface_nat_rule_association#read NetworkInterfaceNatRuleAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

