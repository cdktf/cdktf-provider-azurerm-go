// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob


type ContainerAppJobRegistries struct {
	// The hostname for the Container Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/container_app_job#server ContainerAppJob#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// ID of the System or User Managed Identity used to pull images from the Container Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/container_app_job#identity ContainerAppJob#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// The name of the Secret Reference containing the password value for this user on the Container Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/container_app_job#password_secret_name ContainerAppJob#password_secret_name}
	PasswordSecretName *string `field:"optional" json:"passwordSecretName" yaml:"passwordSecretName"`
	// The username to use for this Container Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/container_app_job#username ContainerAppJob#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

