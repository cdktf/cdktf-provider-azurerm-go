// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistry


type ContainerRegistryIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/container_registry#type ContainerRegistry#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/container_registry#identity_ids ContainerRegistry#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

