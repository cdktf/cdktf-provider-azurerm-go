// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewall


type FirewallVirtualHub struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/firewall#virtual_hub_id Firewall#virtual_hub_id}.
	VirtualHubId *string `field:"required" json:"virtualHubId" yaml:"virtualHubId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/firewall#public_ip_count Firewall#public_ip_count}.
	PublicIpCount *float64 `field:"optional" json:"publicIpCount" yaml:"publicIpCount"`
}

