package automationsoftwareupdateconfiguration


type AutomationSoftwareUpdateConfigurationTarget struct {
	// azure_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/automation_software_update_configuration#azure_query AutomationSoftwareUpdateConfiguration#azure_query}
	AzureQuery interface{} `field:"optional" json:"azureQuery" yaml:"azureQuery"`
	// non_azure_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/automation_software_update_configuration#non_azure_query AutomationSoftwareUpdateConfiguration#non_azure_query}
	NonAzureQuery interface{} `field:"optional" json:"nonAzureQuery" yaml:"nonAzureQuery"`
}

