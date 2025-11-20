// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppTemplate struct {
	// container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/container_app#container ContainerApp#container}
	Container interface{} `field:"required" json:"container" yaml:"container"`
	// azure_queue_scale_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/container_app#azure_queue_scale_rule ContainerApp#azure_queue_scale_rule}
	AzureQueueScaleRule interface{} `field:"optional" json:"azureQueueScaleRule" yaml:"azureQueueScaleRule"`
	// The number of seconds to wait before scaling down the number of instances again.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/container_app#cooldown_period_in_seconds ContainerApp#cooldown_period_in_seconds}
	CooldownPeriodInSeconds *float64 `field:"optional" json:"cooldownPeriodInSeconds" yaml:"cooldownPeriodInSeconds"`
	// custom_scale_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/container_app#custom_scale_rule ContainerApp#custom_scale_rule}
	CustomScaleRule interface{} `field:"optional" json:"customScaleRule" yaml:"customScaleRule"`
	// http_scale_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/container_app#http_scale_rule ContainerApp#http_scale_rule}
	HttpScaleRule interface{} `field:"optional" json:"httpScaleRule" yaml:"httpScaleRule"`
	// init_container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/container_app#init_container ContainerApp#init_container}
	InitContainer interface{} `field:"optional" json:"initContainer" yaml:"initContainer"`
	// The maximum number of replicas for this container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/container_app#max_replicas ContainerApp#max_replicas}
	MaxReplicas *float64 `field:"optional" json:"maxReplicas" yaml:"maxReplicas"`
	// The minimum number of replicas for this container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/container_app#min_replicas ContainerApp#min_replicas}
	MinReplicas *float64 `field:"optional" json:"minReplicas" yaml:"minReplicas"`
	// The interval in seconds used for polling KEDA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/container_app#polling_interval_in_seconds ContainerApp#polling_interval_in_seconds}
	PollingIntervalInSeconds *float64 `field:"optional" json:"pollingIntervalInSeconds" yaml:"pollingIntervalInSeconds"`
	// The suffix for the revision.
	//
	// This value must be unique for the lifetime of the Resource. If omitted the service will use a hash function to create one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/container_app#revision_suffix ContainerApp#revision_suffix}
	RevisionSuffix *string `field:"optional" json:"revisionSuffix" yaml:"revisionSuffix"`
	// tcp_scale_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/container_app#tcp_scale_rule ContainerApp#tcp_scale_rule}
	TcpScaleRule interface{} `field:"optional" json:"tcpScaleRule" yaml:"tcpScaleRule"`
	// The time in seconds after the container is sent the termination signal before the process if forcibly killed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/container_app#termination_grace_period_seconds ContainerApp#termination_grace_period_seconds}
	TerminationGracePeriodSeconds *float64 `field:"optional" json:"terminationGracePeriodSeconds" yaml:"terminationGracePeriodSeconds"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/container_app#volume ContainerApp#volume}
	Volume interface{} `field:"optional" json:"volume" yaml:"volume"`
}

