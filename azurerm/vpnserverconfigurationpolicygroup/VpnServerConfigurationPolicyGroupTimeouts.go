// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnserverconfigurationpolicygroup


type VpnServerConfigurationPolicyGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/vpn_server_configuration_policy_group#create VpnServerConfigurationPolicyGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/vpn_server_configuration_policy_group#delete VpnServerConfigurationPolicyGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/vpn_server_configuration_policy_group#read VpnServerConfigurationPolicyGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/vpn_server_configuration_policy_group#update VpnServerConfigurationPolicyGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

