// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorautoscalesetting


type MonitorAutoscaleSettingProfileCapacity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/monitor_autoscale_setting#default MonitorAutoscaleSetting#default}.
	Default *float64 `field:"required" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/monitor_autoscale_setting#maximum MonitorAutoscaleSetting#maximum}.
	Maximum *float64 `field:"required" json:"maximum" yaml:"maximum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/monitor_autoscale_setting#minimum MonitorAutoscaleSetting#minimum}.
	Minimum *float64 `field:"required" json:"minimum" yaml:"minimum"`
}

