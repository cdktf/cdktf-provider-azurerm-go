package monitoraaddiagnosticsetting


type MonitorAadDiagnosticSettingLog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/monitor_aad_diagnostic_setting#category MonitorAadDiagnosticSetting#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
	// retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/monitor_aad_diagnostic_setting#retention_policy MonitorAadDiagnosticSetting#retention_policy}
	RetentionPolicy *MonitorAadDiagnosticSettingLogRetentionPolicy `field:"required" json:"retentionPolicy" yaml:"retentionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/monitor_aad_diagnostic_setting#enabled MonitorAadDiagnosticSetting#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

