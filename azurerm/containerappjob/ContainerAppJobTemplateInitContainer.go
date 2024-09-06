// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob


type ContainerAppJobTemplateInitContainer struct {
	// The image to use to create the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/container_app_job#image ContainerAppJob#image}
	Image *string `field:"required" json:"image" yaml:"image"`
	// The name of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/container_app_job#name ContainerAppJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of args to pass to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/container_app_job#args ContainerAppJob#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// A command to pass to the container to override the default.
	//
	// This is provided as a list of command line elements without spaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/container_app_job#command ContainerAppJob#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The amount of vCPU to allocate to the container.
	//
	// Possible values include `0.25`, `0.5`, `0.75`, `1.0`, `1.25`, `1.5`, `1.75`, and `2.0`. **NOTE:** `cpu` and `memory` must be specified in `0.25'/'0.5Gi` combination increments. e.g. `1.0` / `2.0` or `0.5` / `1.0`. When there's a workload profile specified, there's no such constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/container_app_job#cpu ContainerAppJob#cpu}
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/container_app_job#env ContainerAppJob#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// The amount of memory to allocate to the container.
	//
	// Possible values include `0.5Gi`, `1.0Gi`, `1.5Gi`, `2.0Gi`, `2.5Gi`, `3.0Gi`, `3.5Gi`, and `4.0Gi`. **NOTE:** `cpu` and `memory` must be specified in `0.25'/'0.5Gi` combination increments. e.g. `1.25` / `2.5Gi` or `0.75` / `1.5Gi`. When there's a workload profile specified, there's no such constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/container_app_job#memory ContainerAppJob#memory}
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// volume_mounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/container_app_job#volume_mounts ContainerAppJob#volume_mounts}
	VolumeMounts interface{} `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
}

