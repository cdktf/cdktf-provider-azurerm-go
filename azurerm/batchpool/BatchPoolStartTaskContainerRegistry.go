// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool


type BatchPoolStartTaskContainerRegistry struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/batch_pool#registry_server BatchPool#registry_server}.
	RegistryServer *string `field:"required" json:"registryServer" yaml:"registryServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/batch_pool#password BatchPool#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The User Assigned Identity to use for Container Registry access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/batch_pool#user_assigned_identity_id BatchPool#user_assigned_identity_id}
	UserAssignedIdentityId *string `field:"optional" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/batch_pool#user_name BatchPool#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

