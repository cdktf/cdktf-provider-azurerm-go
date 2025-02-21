// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgateway


type VirtualNetworkGatewayVpnClientConfigurationIpsecPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/virtual_network_gateway#dh_group VirtualNetworkGateway#dh_group}.
	DhGroup *string `field:"required" json:"dhGroup" yaml:"dhGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/virtual_network_gateway#ike_encryption VirtualNetworkGateway#ike_encryption}.
	IkeEncryption *string `field:"required" json:"ikeEncryption" yaml:"ikeEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/virtual_network_gateway#ike_integrity VirtualNetworkGateway#ike_integrity}.
	IkeIntegrity *string `field:"required" json:"ikeIntegrity" yaml:"ikeIntegrity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/virtual_network_gateway#ipsec_encryption VirtualNetworkGateway#ipsec_encryption}.
	IpsecEncryption *string `field:"required" json:"ipsecEncryption" yaml:"ipsecEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/virtual_network_gateway#ipsec_integrity VirtualNetworkGateway#ipsec_integrity}.
	IpsecIntegrity *string `field:"required" json:"ipsecIntegrity" yaml:"ipsecIntegrity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/virtual_network_gateway#pfs_group VirtualNetworkGateway#pfs_group}.
	PfsGroup *string `field:"required" json:"pfsGroup" yaml:"pfsGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/virtual_network_gateway#sa_data_size_in_kilobytes VirtualNetworkGateway#sa_data_size_in_kilobytes}.
	SaDataSizeInKilobytes *float64 `field:"required" json:"saDataSizeInKilobytes" yaml:"saDataSizeInKilobytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/virtual_network_gateway#sa_lifetime_in_seconds VirtualNetworkGateway#sa_lifetime_in_seconds}.
	SaLifetimeInSeconds *float64 `field:"required" json:"saLifetimeInSeconds" yaml:"saLifetimeInSeconds"`
}

