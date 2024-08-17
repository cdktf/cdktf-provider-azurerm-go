// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotcentralapplicationnetworkruleset


type IotcentralApplicationNetworkRuleSetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/iotcentral_application_network_rule_set#create IotcentralApplicationNetworkRuleSet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/iotcentral_application_network_rule_set#delete IotcentralApplicationNetworkRuleSet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/iotcentral_application_network_rule_set#read IotcentralApplicationNetworkRuleSet#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/iotcentral_application_network_rule_set#update IotcentralApplicationNetworkRuleSet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

