// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob


type ContainerAppJobTemplateContainerVolumeMounts struct {
	// The name of the Volume to be mounted in the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/container_app_job#name ContainerAppJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The path in the container at which to mount this volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/container_app_job#path ContainerAppJob#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The sub path of the volume to be mounted in the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/container_app_job#sub_path ContainerAppJob#sub_path}
	SubPath *string `field:"optional" json:"subPath" yaml:"subPath"`
}

