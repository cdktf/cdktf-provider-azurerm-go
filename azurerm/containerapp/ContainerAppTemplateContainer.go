// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppTemplateContainer struct {
	// The amount of vCPU to allocate to the container.
	//
	// Possible values include `0.25`, `0.5`, `0.75`, `1.0`, `1.25`, `1.5`, `1.75`, and `2.0`. **NOTE:** `cpu` and `memory` must be specified in `0.25'/'0.5Gi` combination increments. e.g. `1.0` / `2.0` or `0.5` / `1.0`. When there's a workload profile specified, there's no such constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/container_app#cpu ContainerApp#cpu}
	Cpu *float64 `field:"required" json:"cpu" yaml:"cpu"`
	// The image to use to create the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/container_app#image ContainerApp#image}
	Image *string `field:"required" json:"image" yaml:"image"`
	// The amount of memory to allocate to the container.
	//
	// Possible values include `0.5Gi`, `1.0Gi`, `1.5Gi`, `2.0Gi`, `2.5Gi`, `3.0Gi`, `3.5Gi`, and `4.0Gi`. **NOTE:** `cpu` and `memory` must be specified in `0.25'/'0.5Gi` combination increments. e.g. `1.25` / `2.5Gi` or `0.75` / `1.5Gi`. When there's a workload profile specified, there's no such constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/container_app#memory ContainerApp#memory}
	Memory *string `field:"required" json:"memory" yaml:"memory"`
	// The name of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/container_app#name ContainerApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of args to pass to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/container_app#args ContainerApp#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// A command to pass to the container to override the default.
	//
	// This is provided as a list of command line elements without spaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/container_app#command ContainerApp#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/container_app#env ContainerApp#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// liveness_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/container_app#liveness_probe ContainerApp#liveness_probe}
	LivenessProbe interface{} `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	// readiness_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/container_app#readiness_probe ContainerApp#readiness_probe}
	ReadinessProbe interface{} `field:"optional" json:"readinessProbe" yaml:"readinessProbe"`
	// startup_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/container_app#startup_probe ContainerApp#startup_probe}
	StartupProbe interface{} `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	// volume_mounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/container_app#volume_mounts ContainerApp#volume_mounts}
	VolumeMounts interface{} `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
}

