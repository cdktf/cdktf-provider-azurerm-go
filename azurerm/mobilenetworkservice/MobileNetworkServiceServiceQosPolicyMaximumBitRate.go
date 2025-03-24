// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkservice


type MobileNetworkServiceServiceQosPolicyMaximumBitRate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/mobile_network_service#downlink MobileNetworkService#downlink}.
	Downlink *string `field:"required" json:"downlink" yaml:"downlink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/mobile_network_service#uplink MobileNetworkService#uplink}.
	Uplink *string `field:"required" json:"uplink" yaml:"uplink"`
}

