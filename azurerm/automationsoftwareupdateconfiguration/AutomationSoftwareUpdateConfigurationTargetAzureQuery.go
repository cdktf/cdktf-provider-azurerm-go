package automationsoftwareupdateconfiguration


type AutomationSoftwareUpdateConfigurationTargetAzureQuery struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#locations AutomationSoftwareUpdateConfiguration#locations}.
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#scope AutomationSoftwareUpdateConfiguration#scope}.
	Scope *[]*string `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#tag_filter AutomationSoftwareUpdateConfiguration#tag_filter}.
	TagFilter *string `field:"optional" json:"tagFilter" yaml:"tagFilter"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#tags AutomationSoftwareUpdateConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}
