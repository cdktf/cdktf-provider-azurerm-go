// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappenvironmentcustomdomain


type ContainerAppEnvironmentCustomDomainTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/container_app_environment_custom_domain#create ContainerAppEnvironmentCustomDomain#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/container_app_environment_custom_domain#delete ContainerAppEnvironmentCustomDomain#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/container_app_environment_custom_domain#read ContainerAppEnvironmentCustomDomain#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/container_app_environment_custom_domain#update ContainerAppEnvironmentCustomDomain#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

