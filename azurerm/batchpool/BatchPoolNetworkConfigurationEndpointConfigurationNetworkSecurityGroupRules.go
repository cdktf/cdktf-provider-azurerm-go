// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool


type BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/batch_pool#access BatchPool#access}.
	Access *string `field:"required" json:"access" yaml:"access"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/batch_pool#priority BatchPool#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/batch_pool#source_address_prefix BatchPool#source_address_prefix}.
	SourceAddressPrefix *string `field:"required" json:"sourceAddressPrefix" yaml:"sourceAddressPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/batch_pool#source_port_ranges BatchPool#source_port_ranges}.
	SourcePortRanges *[]*string `field:"optional" json:"sourcePortRanges" yaml:"sourcePortRanges"`
}

