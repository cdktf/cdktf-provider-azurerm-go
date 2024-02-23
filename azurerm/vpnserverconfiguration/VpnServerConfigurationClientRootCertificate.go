// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnserverconfiguration


type VpnServerConfigurationClientRootCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/vpn_server_configuration#name VpnServerConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/vpn_server_configuration#public_cert_data VpnServerConfiguration#public_cert_data}.
	PublicCertData *string `field:"required" json:"publicCertData" yaml:"publicCertData"`
}

