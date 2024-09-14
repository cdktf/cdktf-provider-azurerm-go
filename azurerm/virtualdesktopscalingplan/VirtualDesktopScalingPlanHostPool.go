// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualdesktopscalingplan


type VirtualDesktopScalingPlanHostPool struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/virtual_desktop_scaling_plan#hostpool_id VirtualDesktopScalingPlan#hostpool_id}.
	HostpoolId *string `field:"required" json:"hostpoolId" yaml:"hostpoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/virtual_desktop_scaling_plan#scaling_plan_enabled VirtualDesktopScalingPlan#scaling_plan_enabled}.
	ScalingPlanEnabled interface{} `field:"required" json:"scalingPlanEnabled" yaml:"scalingPlanEnabled"`
}

