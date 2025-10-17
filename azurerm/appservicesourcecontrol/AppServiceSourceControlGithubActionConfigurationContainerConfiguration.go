// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appservicesourcecontrol


type AppServiceSourceControlGithubActionConfigurationContainerConfiguration struct {
	// The image name for the build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/app_service_source_control#image_name AppServiceSourceControlA#image_name}
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The server URL for the container registry where the build will be hosted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/app_service_source_control#registry_url AppServiceSourceControlA#registry_url}
	RegistryUrl *string `field:"required" json:"registryUrl" yaml:"registryUrl"`
	// The password used to upload the image to the container registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/app_service_source_control#registry_password AppServiceSourceControlA#registry_password}
	RegistryPassword *string `field:"optional" json:"registryPassword" yaml:"registryPassword"`
	// The username used to upload the image to the container registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/app_service_source_control#registry_username AppServiceSourceControlA#registry_username}
	RegistryUsername *string `field:"optional" json:"registryUsername" yaml:"registryUsername"`
}

