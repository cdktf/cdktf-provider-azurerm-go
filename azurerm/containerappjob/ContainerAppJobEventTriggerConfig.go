// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob


type ContainerAppJobEventTriggerConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/container_app_job#parallelism ContainerAppJob#parallelism}.
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/container_app_job#replica_completion_count ContainerAppJob#replica_completion_count}.
	ReplicaCompletionCount *float64 `field:"optional" json:"replicaCompletionCount" yaml:"replicaCompletionCount"`
	// scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/container_app_job#scale ContainerAppJob#scale}
	Scale interface{} `field:"optional" json:"scale" yaml:"scale"`
}

