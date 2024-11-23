// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrytask

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerRegistryTaskConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#container_registry_id ContainerRegistryTask#container_registry_id}.
	ContainerRegistryId *string `field:"required" json:"containerRegistryId" yaml:"containerRegistryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#name ContainerRegistryTask#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#agent_pool_name ContainerRegistryTask#agent_pool_name}.
	AgentPoolName *string `field:"optional" json:"agentPoolName" yaml:"agentPoolName"`
	// agent_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#agent_setting ContainerRegistryTask#agent_setting}
	AgentSetting *ContainerRegistryTaskAgentSetting `field:"optional" json:"agentSetting" yaml:"agentSetting"`
	// base_image_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#base_image_trigger ContainerRegistryTask#base_image_trigger}
	BaseImageTrigger *ContainerRegistryTaskBaseImageTrigger `field:"optional" json:"baseImageTrigger" yaml:"baseImageTrigger"`
	// docker_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#docker_step ContainerRegistryTask#docker_step}
	DockerStep *ContainerRegistryTaskDockerStep `field:"optional" json:"dockerStep" yaml:"dockerStep"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#enabled ContainerRegistryTask#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// encoded_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#encoded_step ContainerRegistryTask#encoded_step}
	EncodedStep *ContainerRegistryTaskEncodedStep `field:"optional" json:"encodedStep" yaml:"encodedStep"`
	// file_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#file_step ContainerRegistryTask#file_step}
	FileStep *ContainerRegistryTaskFileStep `field:"optional" json:"fileStep" yaml:"fileStep"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#id ContainerRegistryTask#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#identity ContainerRegistryTask#identity}
	Identity *ContainerRegistryTaskIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#is_system_task ContainerRegistryTask#is_system_task}.
	IsSystemTask interface{} `field:"optional" json:"isSystemTask" yaml:"isSystemTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#log_template ContainerRegistryTask#log_template}.
	LogTemplate *string `field:"optional" json:"logTemplate" yaml:"logTemplate"`
	// platform block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#platform ContainerRegistryTask#platform}
	Platform *ContainerRegistryTaskPlatform `field:"optional" json:"platform" yaml:"platform"`
	// registry_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#registry_credential ContainerRegistryTask#registry_credential}
	RegistryCredential *ContainerRegistryTaskRegistryCredential `field:"optional" json:"registryCredential" yaml:"registryCredential"`
	// source_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#source_trigger ContainerRegistryTask#source_trigger}
	SourceTrigger interface{} `field:"optional" json:"sourceTrigger" yaml:"sourceTrigger"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#tags ContainerRegistryTask#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#timeout_in_seconds ContainerRegistryTask#timeout_in_seconds}.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#timeouts ContainerRegistryTask#timeouts}
	Timeouts *ContainerRegistryTaskTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// timer_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_task#timer_trigger ContainerRegistryTask#timer_trigger}
	TimerTrigger interface{} `field:"optional" json:"timerTrigger" yaml:"timerTrigger"`
}

