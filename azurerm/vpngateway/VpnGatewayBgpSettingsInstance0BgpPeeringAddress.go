// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpngateway


type VpnGatewayBgpSettingsInstance0BgpPeeringAddress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/vpn_gateway#custom_ips VpnGateway#custom_ips}.
	CustomIps *[]*string `field:"required" json:"customIps" yaml:"customIps"`
}

