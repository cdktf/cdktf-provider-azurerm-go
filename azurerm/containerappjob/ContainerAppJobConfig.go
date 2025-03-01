// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAppJobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#container_app_environment_id ContainerAppJob#container_app_environment_id}.
	ContainerAppEnvironmentId *string `field:"required" json:"containerAppEnvironmentId" yaml:"containerAppEnvironmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#location ContainerAppJob#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#name ContainerAppJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#replica_timeout_in_seconds ContainerAppJob#replica_timeout_in_seconds}.
	ReplicaTimeoutInSeconds *float64 `field:"required" json:"replicaTimeoutInSeconds" yaml:"replicaTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#resource_group_name ContainerAppJob#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#template ContainerAppJob#template}
	Template *ContainerAppJobTemplate `field:"required" json:"template" yaml:"template"`
	// event_trigger_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#event_trigger_config ContainerAppJob#event_trigger_config}
	EventTriggerConfig *ContainerAppJobEventTriggerConfig `field:"optional" json:"eventTriggerConfig" yaml:"eventTriggerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#id ContainerAppJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#identity ContainerAppJob#identity}
	Identity *ContainerAppJobIdentity `field:"optional" json:"identity" yaml:"identity"`
	// manual_trigger_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#manual_trigger_config ContainerAppJob#manual_trigger_config}
	ManualTriggerConfig *ContainerAppJobManualTriggerConfig `field:"optional" json:"manualTriggerConfig" yaml:"manualTriggerConfig"`
	// registry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#registry ContainerAppJob#registry}
	Registry interface{} `field:"optional" json:"registry" yaml:"registry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#replica_retry_limit ContainerAppJob#replica_retry_limit}.
	ReplicaRetryLimit *float64 `field:"optional" json:"replicaRetryLimit" yaml:"replicaRetryLimit"`
	// schedule_trigger_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#schedule_trigger_config ContainerAppJob#schedule_trigger_config}
	ScheduleTriggerConfig *ContainerAppJobScheduleTriggerConfig `field:"optional" json:"scheduleTriggerConfig" yaml:"scheduleTriggerConfig"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#secret ContainerAppJob#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#tags ContainerAppJob#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#timeouts ContainerAppJob#timeouts}
	Timeouts *ContainerAppJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/container_app_job#workload_profile_name ContainerAppJob#workload_profile_name}.
	WorkloadProfileName *string `field:"optional" json:"workloadProfileName" yaml:"workloadProfileName"`
}

