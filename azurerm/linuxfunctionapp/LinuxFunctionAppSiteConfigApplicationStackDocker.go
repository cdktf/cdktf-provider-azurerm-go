// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionapp


type LinuxFunctionAppSiteConfigApplicationStackDocker struct {
	// The name of the Docker image to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/linux_function_app#image_name LinuxFunctionApp#image_name}
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The image tag of the image to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/linux_function_app#image_tag LinuxFunctionApp#image_tag}
	ImageTag *string `field:"required" json:"imageTag" yaml:"imageTag"`
	// The URL of the docker registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/linux_function_app#registry_url LinuxFunctionApp#registry_url}
	RegistryUrl *string `field:"required" json:"registryUrl" yaml:"registryUrl"`
	// The password for the account to use to connect to the registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/linux_function_app#registry_password LinuxFunctionApp#registry_password}
	RegistryPassword *string `field:"optional" json:"registryPassword" yaml:"registryPassword"`
	// The username to use for connections to the registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/linux_function_app#registry_username LinuxFunctionApp#registry_username}
	RegistryUsername *string `field:"optional" json:"registryUsername" yaml:"registryUsername"`
}

