// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob


type ContainerAppJobEventTriggerConfigScale struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/container_app_job#max_executions ContainerAppJob#max_executions}.
	MaxExecutions *float64 `field:"optional" json:"maxExecutions" yaml:"maxExecutions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/container_app_job#min_executions ContainerAppJob#min_executions}.
	MinExecutions *float64 `field:"optional" json:"minExecutions" yaml:"minExecutions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/container_app_job#polling_interval_in_seconds ContainerAppJob#polling_interval_in_seconds}.
	PollingIntervalInSeconds *float64 `field:"optional" json:"pollingIntervalInSeconds" yaml:"pollingIntervalInSeconds"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/container_app_job#rules ContainerAppJob#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

