// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotcentralapplicationnetworkruleset


type IotcentralApplicationNetworkRuleSetIpRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/iotcentral_application_network_rule_set#ip_mask IotcentralApplicationNetworkRuleSet#ip_mask}.
	IpMask *string `field:"required" json:"ipMask" yaml:"ipMask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/iotcentral_application_network_rule_set#name IotcentralApplicationNetworkRuleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

