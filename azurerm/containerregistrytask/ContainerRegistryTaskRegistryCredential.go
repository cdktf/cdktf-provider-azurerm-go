// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrytask


type ContainerRegistryTaskRegistryCredential struct {
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/container_registry_task#custom ContainerRegistryTask#custom}
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/container_registry_task#source ContainerRegistryTask#source}
	Source *ContainerRegistryTaskRegistryCredentialSource `field:"optional" json:"source" yaml:"source"`
}

