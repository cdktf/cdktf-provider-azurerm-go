// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkpacketcorecontrolplane


type MobileNetworkPacketCoreControlPlaneLocalDiagnosticsAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/mobile_network_packet_core_control_plane#authentication_type MobileNetworkPacketCoreControlPlane#authentication_type}.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/mobile_network_packet_core_control_plane#https_server_certificate_url MobileNetworkPacketCoreControlPlane#https_server_certificate_url}.
	HttpsServerCertificateUrl *string `field:"optional" json:"httpsServerCertificateUrl" yaml:"httpsServerCertificateUrl"`
}

