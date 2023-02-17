package monitordiagnosticsetting


type MonitorDiagnosticSettingLogRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_diagnostic_setting#enabled MonitorDiagnosticSetting#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_diagnostic_setting#days MonitorDiagnosticSetting#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

