package datadogmonitor


type DatadogMonitorUser struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/datadog_monitor#email DatadogMonitor#email}.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/datadog_monitor#name DatadogMonitor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/datadog_monitor#phone_number DatadogMonitor#phone_number}.
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

