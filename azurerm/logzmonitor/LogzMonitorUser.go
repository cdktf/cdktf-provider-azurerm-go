package logzmonitor


type LogzMonitorUser struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logz_monitor#email LogzMonitor#email}.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logz_monitor#first_name LogzMonitor#first_name}.
	FirstName *string `field:"required" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logz_monitor#last_name LogzMonitor#last_name}.
	LastName *string `field:"required" json:"lastName" yaml:"lastName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logz_monitor#phone_number LogzMonitor#phone_number}.
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
}
