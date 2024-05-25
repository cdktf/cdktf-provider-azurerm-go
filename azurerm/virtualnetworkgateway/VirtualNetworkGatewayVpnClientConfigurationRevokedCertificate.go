// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgateway


type VirtualNetworkGatewayVpnClientConfigurationRevokedCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/virtual_network_gateway#name VirtualNetworkGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/virtual_network_gateway#thumbprint VirtualNetworkGateway#thumbprint}.
	Thumbprint *string `field:"required" json:"thumbprint" yaml:"thumbprint"`
}

