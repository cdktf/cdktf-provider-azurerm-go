// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorautoscalesetting


type MonitorAutoscaleSettingProfile struct {
	// capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/monitor_autoscale_setting#capacity MonitorAutoscaleSetting#capacity}
	Capacity *MonitorAutoscaleSettingProfileCapacity `field:"required" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/monitor_autoscale_setting#name MonitorAutoscaleSetting#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// fixed_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/monitor_autoscale_setting#fixed_date MonitorAutoscaleSetting#fixed_date}
	FixedDate *MonitorAutoscaleSettingProfileFixedDate `field:"optional" json:"fixedDate" yaml:"fixedDate"`
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/monitor_autoscale_setting#recurrence MonitorAutoscaleSetting#recurrence}
	Recurrence *MonitorAutoscaleSettingProfileRecurrence `field:"optional" json:"recurrence" yaml:"recurrence"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/monitor_autoscale_setting#rule MonitorAutoscaleSetting#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
}

