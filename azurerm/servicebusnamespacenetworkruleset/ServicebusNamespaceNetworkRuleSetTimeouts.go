// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicebusnamespacenetworkruleset


type ServicebusNamespaceNetworkRuleSetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/servicebus_namespace_network_rule_set#create ServicebusNamespaceNetworkRuleSetA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/servicebus_namespace_network_rule_set#delete ServicebusNamespaceNetworkRuleSetA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/servicebus_namespace_network_rule_set#read ServicebusNamespaceNetworkRuleSetA#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/servicebus_namespace_network_rule_set#update ServicebusNamespaceNetworkRuleSetA#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

