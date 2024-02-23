// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanageradminrule


type NetworkManagerAdminRuleSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/network_manager_admin_rule#address_prefix NetworkManagerAdminRule#address_prefix}.
	AddressPrefix *string `field:"required" json:"addressPrefix" yaml:"addressPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/network_manager_admin_rule#address_prefix_type NetworkManagerAdminRule#address_prefix_type}.
	AddressPrefixType *string `field:"required" json:"addressPrefixType" yaml:"addressPrefixType"`
}

