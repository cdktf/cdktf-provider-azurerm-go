package provider


type AzurermProviderFeatures struct {
	// api_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs#api_management AzurermProvider#api_management}
	ApiManagement *AzurermProviderFeaturesApiManagement `field:"optional" json:"apiManagement" yaml:"apiManagement"`
	// app_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs#app_configuration AzurermProvider#app_configuration}
	AppConfiguration *AzurermProviderFeaturesAppConfiguration `field:"optional" json:"appConfiguration" yaml:"appConfiguration"`
	// application_insights block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs#application_insights AzurermProvider#application_insights}
	ApplicationInsights *AzurermProviderFeaturesApplicationInsights `field:"optional" json:"applicationInsights" yaml:"applicationInsights"`
	// cognitive_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs#cognitive_account AzurermProvider#cognitive_account}
	CognitiveAccount *AzurermProviderFeaturesCognitiveAccount `field:"optional" json:"cognitiveAccount" yaml:"cognitiveAccount"`
	// key_vault block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs#key_vault AzurermProvider#key_vault}
	KeyVault *AzurermProviderFeaturesKeyVault `field:"optional" json:"keyVault" yaml:"keyVault"`
	// log_analytics_workspace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs#log_analytics_workspace AzurermProvider#log_analytics_workspace}
	LogAnalyticsWorkspace *AzurermProviderFeaturesLogAnalyticsWorkspace `field:"optional" json:"logAnalyticsWorkspace" yaml:"logAnalyticsWorkspace"`
	// managed_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs#managed_disk AzurermProvider#managed_disk}
	ManagedDisk *AzurermProviderFeaturesManagedDisk `field:"optional" json:"managedDisk" yaml:"managedDisk"`
	// resource_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs#resource_group AzurermProvider#resource_group}
	ResourceGroup *AzurermProviderFeaturesResourceGroup `field:"optional" json:"resourceGroup" yaml:"resourceGroup"`
	// template_deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs#template_deployment AzurermProvider#template_deployment}
	TemplateDeployment *AzurermProviderFeaturesTemplateDeployment `field:"optional" json:"templateDeployment" yaml:"templateDeployment"`
	// virtual_machine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs#virtual_machine AzurermProvider#virtual_machine}
	VirtualMachine *AzurermProviderFeaturesVirtualMachine `field:"optional" json:"virtualMachine" yaml:"virtualMachine"`
	// virtual_machine_scale_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs#virtual_machine_scale_set AzurermProvider#virtual_machine_scale_set}
	VirtualMachineScaleSet *AzurermProviderFeaturesVirtualMachineScaleSet `field:"optional" json:"virtualMachineScaleSet" yaml:"virtualMachineScaleSet"`
}

