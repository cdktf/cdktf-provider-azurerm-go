// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob


type ContainerAppJobTemplateInitContainer struct {
	// The image to use to create the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/container_app_job#image ContainerAppJob#image}
	Image *string `field:"required" json:"image" yaml:"image"`
	// The name of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/container_app_job#name ContainerAppJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of args to pass to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/container_app_job#args ContainerAppJob#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// A command to pass to the container to override the default.
	//
	// This is provided as a list of command line elements without spaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/container_app_job#command ContainerAppJob#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The amount of vCPU to allocate to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/container_app_job#cpu ContainerAppJob#cpu}
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/container_app_job#env ContainerAppJob#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// The amount of memory to allocate to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/container_app_job#memory ContainerAppJob#memory}
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// volume_mounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/container_app_job#volume_mounts ContainerAppJob#volume_mounts}
	VolumeMounts interface{} `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
}

