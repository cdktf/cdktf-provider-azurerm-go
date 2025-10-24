// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallpolicy


type FirewallPolicyDns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/firewall_policy#proxy_enabled FirewallPolicy#proxy_enabled}.
	ProxyEnabled interface{} `field:"optional" json:"proxyEnabled" yaml:"proxyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/firewall_policy#servers FirewallPolicy#servers}.
	Servers *[]*string `field:"optional" json:"servers" yaml:"servers"`
}

