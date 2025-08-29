// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordiagnosticsetting


type MonitorDiagnosticSettingEnabledMetric struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/monitor_diagnostic_setting#category MonitorDiagnosticSetting#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
}

