// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapseworkspace


type SynapseWorkspaceGithubRepo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/synapse_workspace#account_name SynapseWorkspace#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/synapse_workspace#branch_name SynapseWorkspace#branch_name}.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/synapse_workspace#repository_name SynapseWorkspace#repository_name}.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/synapse_workspace#root_folder SynapseWorkspace#root_folder}.
	RootFolder *string `field:"required" json:"rootFolder" yaml:"rootFolder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/synapse_workspace#git_url SynapseWorkspace#git_url}.
	GitUrl *string `field:"optional" json:"gitUrl" yaml:"gitUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/synapse_workspace#last_commit_id SynapseWorkspace#last_commit_id}.
	LastCommitId *string `field:"optional" json:"lastCommitId" yaml:"lastCommitId"`
}

