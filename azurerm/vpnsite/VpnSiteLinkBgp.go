// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnsite


type VpnSiteLinkBgp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/vpn_site#asn VpnSite#asn}.
	Asn *float64 `field:"required" json:"asn" yaml:"asn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/vpn_site#peering_address VpnSite#peering_address}.
	PeeringAddress *string `field:"required" json:"peeringAddress" yaml:"peeringAddress"`
}

