// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob


type ContainerAppJobTemplate struct {
	// container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/container_app_job#container ContainerAppJob#container}
	Container interface{} `field:"required" json:"container" yaml:"container"`
	// init_container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/container_app_job#init_container ContainerAppJob#init_container}
	InitContainer interface{} `field:"optional" json:"initContainer" yaml:"initContainer"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/container_app_job#volume ContainerAppJob#volume}
	Volume interface{} `field:"optional" json:"volume" yaml:"volume"`
}

