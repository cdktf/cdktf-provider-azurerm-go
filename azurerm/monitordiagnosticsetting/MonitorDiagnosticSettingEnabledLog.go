// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordiagnosticsetting


type MonitorDiagnosticSettingEnabledLog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/monitor_diagnostic_setting#category MonitorDiagnosticSetting#category}.
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/monitor_diagnostic_setting#category_group MonitorDiagnosticSetting#category_group}.
	CategoryGroup *string `field:"optional" json:"categoryGroup" yaml:"categoryGroup"`
	// retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/monitor_diagnostic_setting#retention_policy MonitorDiagnosticSetting#retention_policy}
	RetentionPolicy *MonitorDiagnosticSettingEnabledLogRetentionPolicy `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
}

