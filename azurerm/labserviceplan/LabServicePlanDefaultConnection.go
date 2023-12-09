// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package labserviceplan


type LabServicePlanDefaultConnection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/lab_service_plan#client_rdp_access LabServicePlan#client_rdp_access}.
	ClientRdpAccess *string `field:"optional" json:"clientRdpAccess" yaml:"clientRdpAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/lab_service_plan#client_ssh_access LabServicePlan#client_ssh_access}.
	ClientSshAccess *string `field:"optional" json:"clientSshAccess" yaml:"clientSshAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/lab_service_plan#web_rdp_access LabServicePlan#web_rdp_access}.
	WebRdpAccess *string `field:"optional" json:"webRdpAccess" yaml:"webRdpAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/lab_service_plan#web_ssh_access LabServicePlan#web_ssh_access}.
	WebSshAccess *string `field:"optional" json:"webSshAccess" yaml:"webSshAccess"`
}

