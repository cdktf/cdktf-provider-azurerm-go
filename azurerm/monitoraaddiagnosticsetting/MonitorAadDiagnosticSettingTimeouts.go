package monitoraaddiagnosticsetting


type MonitorAadDiagnosticSettingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/monitor_aad_diagnostic_setting#create MonitorAadDiagnosticSetting#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/monitor_aad_diagnostic_setting#delete MonitorAadDiagnosticSetting#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/monitor_aad_diagnostic_setting#read MonitorAadDiagnosticSetting#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/monitor_aad_diagnostic_setting#update MonitorAadDiagnosticSetting#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

