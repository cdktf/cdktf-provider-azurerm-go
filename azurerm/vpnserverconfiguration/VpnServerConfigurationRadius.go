// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnserverconfiguration


type VpnServerConfigurationRadius struct {
	// client_root_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/vpn_server_configuration#client_root_certificate VpnServerConfiguration#client_root_certificate}
	ClientRootCertificate interface{} `field:"optional" json:"clientRootCertificate" yaml:"clientRootCertificate"`
	// server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/vpn_server_configuration#server VpnServerConfiguration#server}
	Server interface{} `field:"optional" json:"server" yaml:"server"`
	// server_root_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/vpn_server_configuration#server_root_certificate VpnServerConfiguration#server_root_certificate}
	ServerRootCertificate interface{} `field:"optional" json:"serverRootCertificate" yaml:"serverRootCertificate"`
}

