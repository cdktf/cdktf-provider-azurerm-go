// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeaturesMachineLearning struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs#purge_soft_deleted_workspace_on_destroy AzurermProvider#purge_soft_deleted_workspace_on_destroy}.
	PurgeSoftDeletedWorkspaceOnDestroy interface{} `field:"optional" json:"purgeSoftDeletedWorkspaceOnDestroy" yaml:"purgeSoftDeletedWorkspaceOnDestroy"`
}

