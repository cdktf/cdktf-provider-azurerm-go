// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppTemplateContainerLivenessProbeHeader struct {
	// The HTTP Header Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/container_app#name ContainerApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The HTTP Header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/container_app#value ContainerApp#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

