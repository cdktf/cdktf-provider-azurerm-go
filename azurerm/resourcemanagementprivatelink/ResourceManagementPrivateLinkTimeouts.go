// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemanagementprivatelink


type ResourceManagementPrivateLinkTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/resource_management_private_link#create ResourceManagementPrivateLink#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/resource_management_private_link#delete ResourceManagementPrivateLink#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/resource_management_private_link#read ResourceManagementPrivateLink#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

