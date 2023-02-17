package datadogmonitor


type DatadogMonitorUser struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor#email DatadogMonitor#email}.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor#name DatadogMonitor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor#phone_number DatadogMonitor#phone_number}.
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

