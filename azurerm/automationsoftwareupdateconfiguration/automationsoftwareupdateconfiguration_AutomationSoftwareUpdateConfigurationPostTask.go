package automationsoftwareupdateconfiguration


type AutomationSoftwareUpdateConfigurationPostTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#parameters AutomationSoftwareUpdateConfiguration#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#source AutomationSoftwareUpdateConfiguration#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

