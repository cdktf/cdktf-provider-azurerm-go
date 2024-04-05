// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorautoscalesetting


type MonitorAutoscaleSettingNotificationWebhook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/monitor_autoscale_setting#service_uri MonitorAutoscaleSetting#service_uri}.
	ServiceUri *string `field:"required" json:"serviceUri" yaml:"serviceUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/monitor_autoscale_setting#properties MonitorAutoscaleSetting#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

