// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnserverconfiguration


type VpnServerConfigurationIpsecPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/vpn_server_configuration#dh_group VpnServerConfiguration#dh_group}.
	DhGroup *string `field:"required" json:"dhGroup" yaml:"dhGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/vpn_server_configuration#ike_encryption VpnServerConfiguration#ike_encryption}.
	IkeEncryption *string `field:"required" json:"ikeEncryption" yaml:"ikeEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/vpn_server_configuration#ike_integrity VpnServerConfiguration#ike_integrity}.
	IkeIntegrity *string `field:"required" json:"ikeIntegrity" yaml:"ikeIntegrity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/vpn_server_configuration#ipsec_encryption VpnServerConfiguration#ipsec_encryption}.
	IpsecEncryption *string `field:"required" json:"ipsecEncryption" yaml:"ipsecEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/vpn_server_configuration#ipsec_integrity VpnServerConfiguration#ipsec_integrity}.
	IpsecIntegrity *string `field:"required" json:"ipsecIntegrity" yaml:"ipsecIntegrity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/vpn_server_configuration#pfs_group VpnServerConfiguration#pfs_group}.
	PfsGroup *string `field:"required" json:"pfsGroup" yaml:"pfsGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/vpn_server_configuration#sa_data_size_kilobytes VpnServerConfiguration#sa_data_size_kilobytes}.
	SaDataSizeKilobytes *float64 `field:"required" json:"saDataSizeKilobytes" yaml:"saDataSizeKilobytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/vpn_server_configuration#sa_lifetime_seconds VpnServerConfiguration#sa_lifetime_seconds}.
	SaLifetimeSeconds *float64 `field:"required" json:"saLifetimeSeconds" yaml:"saLifetimeSeconds"`
}

