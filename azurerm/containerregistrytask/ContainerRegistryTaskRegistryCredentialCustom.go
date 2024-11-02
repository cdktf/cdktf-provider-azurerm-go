// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrytask


type ContainerRegistryTaskRegistryCredentialCustom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/container_registry_task#login_server ContainerRegistryTask#login_server}.
	LoginServer *string `field:"required" json:"loginServer" yaml:"loginServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/container_registry_task#identity ContainerRegistryTask#identity}.
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/container_registry_task#password ContainerRegistryTask#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/container_registry_task#username ContainerRegistryTask#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

