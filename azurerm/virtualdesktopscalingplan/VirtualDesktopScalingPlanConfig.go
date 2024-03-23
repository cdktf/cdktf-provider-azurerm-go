// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualdesktopscalingplan

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualDesktopScalingPlanConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/virtual_desktop_scaling_plan#location VirtualDesktopScalingPlan#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/virtual_desktop_scaling_plan#name VirtualDesktopScalingPlan#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/virtual_desktop_scaling_plan#resource_group_name VirtualDesktopScalingPlan#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/virtual_desktop_scaling_plan#schedule VirtualDesktopScalingPlan#schedule}
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/virtual_desktop_scaling_plan#time_zone VirtualDesktopScalingPlan#time_zone}.
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/virtual_desktop_scaling_plan#description VirtualDesktopScalingPlan#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/virtual_desktop_scaling_plan#exclusion_tag VirtualDesktopScalingPlan#exclusion_tag}.
	ExclusionTag *string `field:"optional" json:"exclusionTag" yaml:"exclusionTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/virtual_desktop_scaling_plan#friendly_name VirtualDesktopScalingPlan#friendly_name}.
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// host_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/virtual_desktop_scaling_plan#host_pool VirtualDesktopScalingPlan#host_pool}
	HostPool interface{} `field:"optional" json:"hostPool" yaml:"hostPool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/virtual_desktop_scaling_plan#id VirtualDesktopScalingPlan#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/virtual_desktop_scaling_plan#tags VirtualDesktopScalingPlan#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/virtual_desktop_scaling_plan#timeouts VirtualDesktopScalingPlan#timeouts}
	Timeouts *VirtualDesktopScalingPlanTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

