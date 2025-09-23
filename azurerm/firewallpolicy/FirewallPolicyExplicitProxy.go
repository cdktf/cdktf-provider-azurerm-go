// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallpolicy


type FirewallPolicyExplicitProxy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/firewall_policy#enabled FirewallPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/firewall_policy#enable_pac_file FirewallPolicy#enable_pac_file}.
	EnablePacFile interface{} `field:"optional" json:"enablePacFile" yaml:"enablePacFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/firewall_policy#http_port FirewallPolicy#http_port}.
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/firewall_policy#https_port FirewallPolicy#https_port}.
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/firewall_policy#pac_file FirewallPolicy#pac_file}.
	PacFile *string `field:"optional" json:"pacFile" yaml:"pacFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/firewall_policy#pac_file_port FirewallPolicy#pac_file_port}.
	PacFilePort *float64 `field:"optional" json:"pacFilePort" yaml:"pacFilePort"`
}

