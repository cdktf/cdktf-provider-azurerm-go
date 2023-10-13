// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorautoscalesetting


type MonitorAutoscaleSettingProfileRule struct {
	// metric_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_autoscale_setting#metric_trigger MonitorAutoscaleSetting#metric_trigger}
	MetricTrigger *MonitorAutoscaleSettingProfileRuleMetricTrigger `field:"required" json:"metricTrigger" yaml:"metricTrigger"`
	// scale_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_autoscale_setting#scale_action MonitorAutoscaleSetting#scale_action}
	ScaleAction *MonitorAutoscaleSettingProfileRuleScaleAction `field:"required" json:"scaleAction" yaml:"scaleAction"`
}

