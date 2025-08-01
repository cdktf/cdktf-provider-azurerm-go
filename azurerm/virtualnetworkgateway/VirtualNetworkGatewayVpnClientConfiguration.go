// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgateway


type VirtualNetworkGatewayVpnClientConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#address_space VirtualNetworkGateway#address_space}.
	AddressSpace *[]*string `field:"required" json:"addressSpace" yaml:"addressSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#aad_audience VirtualNetworkGateway#aad_audience}.
	AadAudience *string `field:"optional" json:"aadAudience" yaml:"aadAudience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#aad_issuer VirtualNetworkGateway#aad_issuer}.
	AadIssuer *string `field:"optional" json:"aadIssuer" yaml:"aadIssuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#aad_tenant VirtualNetworkGateway#aad_tenant}.
	AadTenant *string `field:"optional" json:"aadTenant" yaml:"aadTenant"`
	// ipsec_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#ipsec_policy VirtualNetworkGateway#ipsec_policy}
	IpsecPolicy *VirtualNetworkGatewayVpnClientConfigurationIpsecPolicy `field:"optional" json:"ipsecPolicy" yaml:"ipsecPolicy"`
	// radius_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#radius_server VirtualNetworkGateway#radius_server}
	RadiusServer interface{} `field:"optional" json:"radiusServer" yaml:"radiusServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#radius_server_address VirtualNetworkGateway#radius_server_address}.
	RadiusServerAddress *string `field:"optional" json:"radiusServerAddress" yaml:"radiusServerAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#radius_server_secret VirtualNetworkGateway#radius_server_secret}.
	RadiusServerSecret *string `field:"optional" json:"radiusServerSecret" yaml:"radiusServerSecret"`
	// revoked_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#revoked_certificate VirtualNetworkGateway#revoked_certificate}
	RevokedCertificate interface{} `field:"optional" json:"revokedCertificate" yaml:"revokedCertificate"`
	// root_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#root_certificate VirtualNetworkGateway#root_certificate}
	RootCertificate interface{} `field:"optional" json:"rootCertificate" yaml:"rootCertificate"`
	// virtual_network_gateway_client_connection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#virtual_network_gateway_client_connection VirtualNetworkGateway#virtual_network_gateway_client_connection}
	VirtualNetworkGatewayClientConnection interface{} `field:"optional" json:"virtualNetworkGatewayClientConnection" yaml:"virtualNetworkGatewayClientConnection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#vpn_auth_types VirtualNetworkGateway#vpn_auth_types}.
	VpnAuthTypes *[]*string `field:"optional" json:"vpnAuthTypes" yaml:"vpnAuthTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/virtual_network_gateway#vpn_client_protocols VirtualNetworkGateway#vpn_client_protocols}.
	VpnClientProtocols *[]*string `field:"optional" json:"vpnClientProtocols" yaml:"vpnClientProtocols"`
}

