// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapseworkspaceaadadmin


type SynapseWorkspaceAadAdminTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/synapse_workspace_aad_admin#create SynapseWorkspaceAadAdmin#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/synapse_workspace_aad_admin#delete SynapseWorkspaceAadAdmin#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/synapse_workspace_aad_admin#read SynapseWorkspaceAadAdmin#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/synapse_workspace_aad_admin#update SynapseWorkspaceAadAdmin#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

