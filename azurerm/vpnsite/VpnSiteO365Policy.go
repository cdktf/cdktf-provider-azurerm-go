// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnsite


type VpnSiteO365Policy struct {
	// traffic_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/vpn_site#traffic_category VpnSite#traffic_category}
	TrafficCategory *VpnSiteO365PolicyTrafficCategory `field:"optional" json:"trafficCategory" yaml:"trafficCategory"`
}

