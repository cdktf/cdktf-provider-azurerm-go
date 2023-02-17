package monitordiagnosticsetting


type MonitorDiagnosticSettingMetric struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_diagnostic_setting#category MonitorDiagnosticSetting#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_diagnostic_setting#enabled MonitorDiagnosticSetting#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_diagnostic_setting#retention_policy MonitorDiagnosticSetting#retention_policy}
	RetentionPolicy *MonitorDiagnosticSettingMetricRetentionPolicy `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
}

