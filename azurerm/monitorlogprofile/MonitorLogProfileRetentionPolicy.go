// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorlogprofile


type MonitorLogProfileRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/monitor_log_profile#enabled MonitorLogProfile#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/monitor_log_profile#days MonitorLogProfile#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

