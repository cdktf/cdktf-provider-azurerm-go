// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package labservicelab


type LabServiceLabConnectionSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/lab_service_lab#client_rdp_access LabServiceLab#client_rdp_access}.
	ClientRdpAccess *string `field:"optional" json:"clientRdpAccess" yaml:"clientRdpAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/lab_service_lab#client_ssh_access LabServiceLab#client_ssh_access}.
	ClientSshAccess *string `field:"optional" json:"clientSshAccess" yaml:"clientSshAccess"`
}

