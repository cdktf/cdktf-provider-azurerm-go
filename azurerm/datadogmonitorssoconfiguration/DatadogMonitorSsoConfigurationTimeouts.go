package datadogmonitorssoconfiguration


type DatadogMonitorSsoConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor_sso_configuration#create DatadogMonitorSsoConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor_sso_configuration#delete DatadogMonitorSsoConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor_sso_configuration#read DatadogMonitorSsoConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor_sso_configuration#update DatadogMonitorSsoConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
