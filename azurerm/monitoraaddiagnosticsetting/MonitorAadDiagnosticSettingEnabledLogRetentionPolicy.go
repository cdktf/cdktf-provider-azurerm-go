// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoraaddiagnosticsetting


type MonitorAadDiagnosticSettingEnabledLogRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/monitor_aad_diagnostic_setting#days MonitorAadDiagnosticSetting#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/monitor_aad_diagnostic_setting#enabled MonitorAadDiagnosticSetting#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

