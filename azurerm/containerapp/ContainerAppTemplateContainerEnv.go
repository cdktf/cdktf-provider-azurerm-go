// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppTemplateContainerEnv struct {
	// The name of the environment variable for the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/container_app#name ContainerApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the secret that contains the value for this environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/container_app#secret_name ContainerApp#secret_name}
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
	// The value for this environment variable. **NOTE:** This value is ignored if `secret_name` is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/container_app#value ContainerApp#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

