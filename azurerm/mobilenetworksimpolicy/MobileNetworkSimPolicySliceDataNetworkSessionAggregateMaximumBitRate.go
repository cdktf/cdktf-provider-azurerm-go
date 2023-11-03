// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworksimpolicy


type MobileNetworkSimPolicySliceDataNetworkSessionAggregateMaximumBitRate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/mobile_network_sim_policy#downlink MobileNetworkSimPolicy#downlink}.
	Downlink *string `field:"required" json:"downlink" yaml:"downlink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/mobile_network_sim_policy#uplink MobileNetworkSimPolicy#uplink}.
	Uplink *string `field:"required" json:"uplink" yaml:"uplink"`
}

