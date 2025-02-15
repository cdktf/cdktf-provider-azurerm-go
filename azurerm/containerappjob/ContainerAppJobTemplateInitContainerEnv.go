// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob


type ContainerAppJobTemplateInitContainerEnv struct {
	// The name of the environment variable for the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/container_app_job#name ContainerAppJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the secret that contains the value for this environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/container_app_job#secret_name ContainerAppJob#secret_name}
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
	// The value for this environment variable. **NOTE:** This value is ignored if `secret_name` is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/container_app_job#value ContainerAppJob#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

