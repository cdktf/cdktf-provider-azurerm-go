// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeatures struct {
	// api_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#api_management AzurermProvider#api_management}
	ApiManagement *AzurermProviderFeaturesApiManagement `field:"optional" json:"apiManagement" yaml:"apiManagement"`
	// app_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#app_configuration AzurermProvider#app_configuration}
	AppConfiguration *AzurermProviderFeaturesAppConfiguration `field:"optional" json:"appConfiguration" yaml:"appConfiguration"`
	// application_insights block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#application_insights AzurermProvider#application_insights}
	ApplicationInsights *AzurermProviderFeaturesApplicationInsights `field:"optional" json:"applicationInsights" yaml:"applicationInsights"`
	// cognitive_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#cognitive_account AzurermProvider#cognitive_account}
	CognitiveAccount *AzurermProviderFeaturesCognitiveAccount `field:"optional" json:"cognitiveAccount" yaml:"cognitiveAccount"`
	// key_vault block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#key_vault AzurermProvider#key_vault}
	KeyVault *AzurermProviderFeaturesKeyVault `field:"optional" json:"keyVault" yaml:"keyVault"`
	// log_analytics_workspace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#log_analytics_workspace AzurermProvider#log_analytics_workspace}
	LogAnalyticsWorkspace *AzurermProviderFeaturesLogAnalyticsWorkspace `field:"optional" json:"logAnalyticsWorkspace" yaml:"logAnalyticsWorkspace"`
	// machine_learning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#machine_learning AzurermProvider#machine_learning}
	MachineLearning *AzurermProviderFeaturesMachineLearning `field:"optional" json:"machineLearning" yaml:"machineLearning"`
	// managed_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#managed_disk AzurermProvider#managed_disk}
	ManagedDisk *AzurermProviderFeaturesManagedDisk `field:"optional" json:"managedDisk" yaml:"managedDisk"`
	// postgresql_flexible_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#postgresql_flexible_server AzurermProvider#postgresql_flexible_server}
	PostgresqlFlexibleServer *AzurermProviderFeaturesPostgresqlFlexibleServer `field:"optional" json:"postgresqlFlexibleServer" yaml:"postgresqlFlexibleServer"`
	// recovery_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#recovery_service AzurermProvider#recovery_service}
	RecoveryService *AzurermProviderFeaturesRecoveryService `field:"optional" json:"recoveryService" yaml:"recoveryService"`
	// recovery_services_vaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#recovery_services_vaults AzurermProvider#recovery_services_vaults}
	RecoveryServicesVaults *AzurermProviderFeaturesRecoveryServicesVaults `field:"optional" json:"recoveryServicesVaults" yaml:"recoveryServicesVaults"`
	// resource_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#resource_group AzurermProvider#resource_group}
	ResourceGroup *AzurermProviderFeaturesResourceGroup `field:"optional" json:"resourceGroup" yaml:"resourceGroup"`
	// subscription block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#subscription AzurermProvider#subscription}
	Subscription *AzurermProviderFeaturesSubscription `field:"optional" json:"subscription" yaml:"subscription"`
	// template_deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#template_deployment AzurermProvider#template_deployment}
	TemplateDeployment *AzurermProviderFeaturesTemplateDeployment `field:"optional" json:"templateDeployment" yaml:"templateDeployment"`
	// virtual_machine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#virtual_machine AzurermProvider#virtual_machine}
	VirtualMachine *AzurermProviderFeaturesVirtualMachine `field:"optional" json:"virtualMachine" yaml:"virtualMachine"`
	// virtual_machine_scale_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#virtual_machine_scale_set AzurermProvider#virtual_machine_scale_set}
	VirtualMachineScaleSet *AzurermProviderFeaturesVirtualMachineScaleSet `field:"optional" json:"virtualMachineScaleSet" yaml:"virtualMachineScaleSet"`
}

