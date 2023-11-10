// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapseworkspace


type SynapseWorkspaceAadAdmin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/synapse_workspace#login SynapseWorkspace#login}.
	Login *string `field:"optional" json:"login" yaml:"login"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/synapse_workspace#object_id SynapseWorkspace#object_id}.
	ObjectId *string `field:"optional" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/synapse_workspace#tenant_id SynapseWorkspace#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

