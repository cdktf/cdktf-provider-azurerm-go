// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob


type ContainerAppJobScheduleTriggerConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/container_app_job#cron_expression ContainerAppJob#cron_expression}.
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/container_app_job#parallelism ContainerAppJob#parallelism}.
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/container_app_job#replica_completion_count ContainerAppJob#replica_completion_count}.
	ReplicaCompletionCount *float64 `field:"optional" json:"replicaCompletionCount" yaml:"replicaCompletionCount"`
}

