// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappenvironmentdaprcomponent


type ContainerAppEnvironmentDaprComponentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/container_app_environment_dapr_component#create ContainerAppEnvironmentDaprComponent#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/container_app_environment_dapr_component#delete ContainerAppEnvironmentDaprComponent#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/container_app_environment_dapr_component#read ContainerAppEnvironmentDaprComponent#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/container_app_environment_dapr_component#update ContainerAppEnvironmentDaprComponent#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

