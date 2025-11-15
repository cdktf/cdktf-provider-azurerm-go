// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagerroutingrule


type NetworkManagerRoutingRuleDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/network_manager_routing_rule#address NetworkManagerRoutingRule#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/network_manager_routing_rule#type NetworkManagerRoutingRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

