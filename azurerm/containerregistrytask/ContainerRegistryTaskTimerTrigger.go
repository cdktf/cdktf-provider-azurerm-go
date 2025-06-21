// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrytask


type ContainerRegistryTaskTimerTrigger struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/container_registry_task#name ContainerRegistryTask#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/container_registry_task#schedule ContainerRegistryTask#schedule}.
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/container_registry_task#enabled ContainerRegistryTask#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

