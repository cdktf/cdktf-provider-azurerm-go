// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrytaskschedulerunnow

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerRegistryTaskScheduleRunNowConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.5.0/docs/resources/container_registry_task_schedule_run_now#container_registry_task_id ContainerRegistryTaskScheduleRunNow#container_registry_task_id}.
	ContainerRegistryTaskId *string `field:"required" json:"containerRegistryTaskId" yaml:"containerRegistryTaskId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.5.0/docs/resources/container_registry_task_schedule_run_now#id ContainerRegistryTaskScheduleRunNow#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.5.0/docs/resources/container_registry_task_schedule_run_now#timeouts ContainerRegistryTaskScheduleRunNow#timeouts}
	Timeouts *ContainerRegistryTaskScheduleRunNowTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

