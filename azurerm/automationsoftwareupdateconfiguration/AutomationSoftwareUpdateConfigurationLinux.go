package automationsoftwareupdateconfiguration


type AutomationSoftwareUpdateConfigurationLinux struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/automation_software_update_configuration#classification_included AutomationSoftwareUpdateConfiguration#classification_included}.
	ClassificationIncluded *string `field:"optional" json:"classificationIncluded" yaml:"classificationIncluded"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/automation_software_update_configuration#excluded_packages AutomationSoftwareUpdateConfiguration#excluded_packages}.
	ExcludedPackages *[]*string `field:"optional" json:"excludedPackages" yaml:"excludedPackages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/automation_software_update_configuration#included_packages AutomationSoftwareUpdateConfiguration#included_packages}.
	IncludedPackages *[]*string `field:"optional" json:"includedPackages" yaml:"includedPackages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/automation_software_update_configuration#reboot AutomationSoftwareUpdateConfiguration#reboot}.
	Reboot *string `field:"optional" json:"reboot" yaml:"reboot"`
}

