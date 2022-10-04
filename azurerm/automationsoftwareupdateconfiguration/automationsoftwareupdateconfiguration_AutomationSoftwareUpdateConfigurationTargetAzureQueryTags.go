package automationsoftwareupdateconfiguration


type AutomationSoftwareUpdateConfigurationTargetAzureQueryTags struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#tag AutomationSoftwareUpdateConfiguration#tag}.
	Tag *string `field:"required" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#values AutomationSoftwareUpdateConfiguration#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

