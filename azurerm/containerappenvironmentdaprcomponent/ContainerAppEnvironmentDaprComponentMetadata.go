// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappenvironmentdaprcomponent


type ContainerAppEnvironmentDaprComponentMetadata struct {
	// The name of the Metadata configuration item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/container_app_environment_dapr_component#name ContainerAppEnvironmentDaprComponent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of a secret specified in the `secrets` block that contains the value for this metadata configuration item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/container_app_environment_dapr_component#secret_name ContainerAppEnvironmentDaprComponent#secret_name}
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
	// The value for this metadata configuration item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/container_app_environment_dapr_component#value ContainerAppEnvironmentDaprComponent#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

