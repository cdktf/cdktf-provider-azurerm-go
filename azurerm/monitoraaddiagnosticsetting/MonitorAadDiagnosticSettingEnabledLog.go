// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoraaddiagnosticsetting


type MonitorAadDiagnosticSettingEnabledLog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/monitor_aad_diagnostic_setting#category MonitorAadDiagnosticSetting#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
	// retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/monitor_aad_diagnostic_setting#retention_policy MonitorAadDiagnosticSetting#retention_policy}
	RetentionPolicy *MonitorAadDiagnosticSettingEnabledLogRetentionPolicy `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
}

