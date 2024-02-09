// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapseworkspacesecurityalertpolicy


type SynapseWorkspaceSecurityAlertPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/synapse_workspace_security_alert_policy#create SynapseWorkspaceSecurityAlertPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/synapse_workspace_security_alert_policy#delete SynapseWorkspaceSecurityAlertPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/synapse_workspace_security_alert_policy#read SynapseWorkspaceSecurityAlertPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/synapse_workspace_security_alert_policy#update SynapseWorkspaceSecurityAlertPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

