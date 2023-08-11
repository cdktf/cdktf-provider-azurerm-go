package monitoractiongroup


type MonitorActionGroupVoiceReceiver struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/monitor_action_group#country_code MonitorActionGroup#country_code}.
	CountryCode *string `field:"required" json:"countryCode" yaml:"countryCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/monitor_action_group#name MonitorActionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/monitor_action_group#phone_number MonitorActionGroup#phone_number}.
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
}

