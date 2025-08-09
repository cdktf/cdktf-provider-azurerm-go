// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpngatewayconnection


type VpnGatewayConnectionVpnLinkIpsecPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/vpn_gateway_connection#dh_group VpnGatewayConnection#dh_group}.
	DhGroup *string `field:"required" json:"dhGroup" yaml:"dhGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/vpn_gateway_connection#encryption_algorithm VpnGatewayConnection#encryption_algorithm}.
	EncryptionAlgorithm *string `field:"required" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/vpn_gateway_connection#ike_encryption_algorithm VpnGatewayConnection#ike_encryption_algorithm}.
	IkeEncryptionAlgorithm *string `field:"required" json:"ikeEncryptionAlgorithm" yaml:"ikeEncryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/vpn_gateway_connection#ike_integrity_algorithm VpnGatewayConnection#ike_integrity_algorithm}.
	IkeIntegrityAlgorithm *string `field:"required" json:"ikeIntegrityAlgorithm" yaml:"ikeIntegrityAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/vpn_gateway_connection#integrity_algorithm VpnGatewayConnection#integrity_algorithm}.
	IntegrityAlgorithm *string `field:"required" json:"integrityAlgorithm" yaml:"integrityAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/vpn_gateway_connection#pfs_group VpnGatewayConnection#pfs_group}.
	PfsGroup *string `field:"required" json:"pfsGroup" yaml:"pfsGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/vpn_gateway_connection#sa_data_size_kb VpnGatewayConnection#sa_data_size_kb}.
	SaDataSizeKb *float64 `field:"required" json:"saDataSizeKb" yaml:"saDataSizeKb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/vpn_gateway_connection#sa_lifetime_sec VpnGatewayConnection#sa_lifetime_sec}.
	SaLifetimeSec *float64 `field:"required" json:"saLifetimeSec" yaml:"saLifetimeSec"`
}

