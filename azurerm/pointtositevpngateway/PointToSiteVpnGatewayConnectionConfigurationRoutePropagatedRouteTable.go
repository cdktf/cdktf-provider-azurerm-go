// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pointtositevpngateway


type PointToSiteVpnGatewayConnectionConfigurationRoutePropagatedRouteTable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/point_to_site_vpn_gateway#ids PointToSiteVpnGateway#ids}.
	Ids *[]*string `field:"required" json:"ids" yaml:"ids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/point_to_site_vpn_gateway#labels PointToSiteVpnGateway#labels}.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
}

