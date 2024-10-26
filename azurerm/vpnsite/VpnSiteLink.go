// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnsite


type VpnSiteLink struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/vpn_site#name VpnSite#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// bgp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/vpn_site#bgp VpnSite#bgp}
	Bgp *VpnSiteLinkBgp `field:"optional" json:"bgp" yaml:"bgp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/vpn_site#fqdn VpnSite#fqdn}.
	Fqdn *string `field:"optional" json:"fqdn" yaml:"fqdn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/vpn_site#ip_address VpnSite#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/vpn_site#provider_name VpnSite#provider_name}.
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/vpn_site#speed_in_mbps VpnSite#speed_in_mbps}.
	SpeedInMbps *float64 `field:"optional" json:"speedInMbps" yaml:"speedInMbps"`
}

