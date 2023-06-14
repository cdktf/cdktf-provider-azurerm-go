package automationsoftwareupdateconfiguration


type AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrence struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/automation_software_update_configuration#day AutomationSoftwareUpdateConfiguration#day}.
	Day *string `field:"required" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/automation_software_update_configuration#occurrence AutomationSoftwareUpdateConfiguration#occurrence}.
	Occurrence *float64 `field:"required" json:"occurrence" yaml:"occurrence"`
}

