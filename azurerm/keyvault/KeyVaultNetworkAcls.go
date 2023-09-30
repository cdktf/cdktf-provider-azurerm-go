// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvault


type KeyVaultNetworkAcls struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/key_vault#bypass KeyVault#bypass}.
	Bypass *string `field:"required" json:"bypass" yaml:"bypass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/key_vault#default_action KeyVault#default_action}.
	DefaultAction *string `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/key_vault#ip_rules KeyVault#ip_rules}.
	IpRules *[]*string `field:"optional" json:"ipRules" yaml:"ipRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/key_vault#virtual_network_subnet_ids KeyVault#virtual_network_subnet_ids}.
	VirtualNetworkSubnetIds *[]*string `field:"optional" json:"virtualNetworkSubnetIds" yaml:"virtualNetworkSubnetIds"`
}

