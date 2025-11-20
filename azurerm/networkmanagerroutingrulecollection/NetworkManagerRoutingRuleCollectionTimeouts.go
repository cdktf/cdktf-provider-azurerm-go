// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagerroutingrulecollection


type NetworkManagerRoutingRuleCollectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/network_manager_routing_rule_collection#create NetworkManagerRoutingRuleCollection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/network_manager_routing_rule_collection#delete NetworkManagerRoutingRuleCollection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/network_manager_routing_rule_collection#read NetworkManagerRoutingRuleCollection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/network_manager_routing_rule_collection#update NetworkManagerRoutingRuleCollection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

