// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pointtositevpngateway


type PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/point_to_site_vpn_gateway#address_prefixes PointToSiteVpnGateway#address_prefixes}.
	AddressPrefixes *[]*string `field:"required" json:"addressPrefixes" yaml:"addressPrefixes"`
}

