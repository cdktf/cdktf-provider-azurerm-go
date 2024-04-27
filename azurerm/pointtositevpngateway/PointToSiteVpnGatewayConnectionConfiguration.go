// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pointtositevpngateway


type PointToSiteVpnGatewayConnectionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/point_to_site_vpn_gateway#name PointToSiteVpnGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// vpn_client_address_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/point_to_site_vpn_gateway#vpn_client_address_pool PointToSiteVpnGateway#vpn_client_address_pool}
	VpnClientAddressPool *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool `field:"required" json:"vpnClientAddressPool" yaml:"vpnClientAddressPool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/point_to_site_vpn_gateway#internet_security_enabled PointToSiteVpnGateway#internet_security_enabled}.
	InternetSecurityEnabled interface{} `field:"optional" json:"internetSecurityEnabled" yaml:"internetSecurityEnabled"`
	// route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.101.0/docs/resources/point_to_site_vpn_gateway#route PointToSiteVpnGateway#route}
	Route *PointToSiteVpnGatewayConnectionConfigurationRoute `field:"optional" json:"route" yaml:"route"`
}

