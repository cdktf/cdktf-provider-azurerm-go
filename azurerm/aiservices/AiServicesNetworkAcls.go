// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aiservices


type AiServicesNetworkAcls struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/ai_services#default_action AiServices#default_action}.
	DefaultAction *string `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/ai_services#bypass AiServices#bypass}.
	Bypass *string `field:"optional" json:"bypass" yaml:"bypass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/ai_services#ip_rules AiServices#ip_rules}.
	IpRules *[]*string `field:"optional" json:"ipRules" yaml:"ipRules"`
	// virtual_network_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/ai_services#virtual_network_rules AiServices#virtual_network_rules}
	VirtualNetworkRules interface{} `field:"optional" json:"virtualNetworkRules" yaml:"virtualNetworkRules"`
}

