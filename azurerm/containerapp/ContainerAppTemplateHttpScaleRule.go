// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppTemplateHttpScaleRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/container_app#concurrent_requests ContainerApp#concurrent_requests}.
	ConcurrentRequests *string `field:"required" json:"concurrentRequests" yaml:"concurrentRequests"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/container_app#name ContainerApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/container_app#authentication ContainerApp#authentication}
	Authentication interface{} `field:"optional" json:"authentication" yaml:"authentication"`
}

