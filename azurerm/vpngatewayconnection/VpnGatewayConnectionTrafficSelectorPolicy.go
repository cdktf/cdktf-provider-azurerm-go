// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpngatewayconnection


type VpnGatewayConnectionTrafficSelectorPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/vpn_gateway_connection#local_address_ranges VpnGatewayConnection#local_address_ranges}.
	LocalAddressRanges *[]*string `field:"required" json:"localAddressRanges" yaml:"localAddressRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/vpn_gateway_connection#remote_address_ranges VpnGatewayConnection#remote_address_ranges}.
	RemoteAddressRanges *[]*string `field:"required" json:"remoteAddressRanges" yaml:"remoteAddressRanges"`
}

