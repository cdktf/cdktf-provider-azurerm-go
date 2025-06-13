// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppTemplateAzureQueueScaleRule struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/container_app#authentication ContainerApp#authentication}
	Authentication interface{} `field:"required" json:"authentication" yaml:"authentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/container_app#name ContainerApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/container_app#queue_length ContainerApp#queue_length}.
	QueueLength *float64 `field:"required" json:"queueLength" yaml:"queueLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/container_app#queue_name ContainerApp#queue_name}.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
}

