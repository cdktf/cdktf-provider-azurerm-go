// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeatures struct {
	// api_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#api_management AzurermProvider#api_management}
	ApiManagement interface{} `field:"optional" json:"apiManagement" yaml:"apiManagement"`
	// app_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#app_configuration AzurermProvider#app_configuration}
	AppConfiguration interface{} `field:"optional" json:"appConfiguration" yaml:"appConfiguration"`
	// application_insights block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#application_insights AzurermProvider#application_insights}
	ApplicationInsights interface{} `field:"optional" json:"applicationInsights" yaml:"applicationInsights"`
	// cognitive_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#cognitive_account AzurermProvider#cognitive_account}
	CognitiveAccount interface{} `field:"optional" json:"cognitiveAccount" yaml:"cognitiveAccount"`
	// key_vault block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#key_vault AzurermProvider#key_vault}
	KeyVault interface{} `field:"optional" json:"keyVault" yaml:"keyVault"`
	// log_analytics_workspace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#log_analytics_workspace AzurermProvider#log_analytics_workspace}
	LogAnalyticsWorkspace interface{} `field:"optional" json:"logAnalyticsWorkspace" yaml:"logAnalyticsWorkspace"`
	// machine_learning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#machine_learning AzurermProvider#machine_learning}
	MachineLearning interface{} `field:"optional" json:"machineLearning" yaml:"machineLearning"`
	// managed_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#managed_disk AzurermProvider#managed_disk}
	ManagedDisk interface{} `field:"optional" json:"managedDisk" yaml:"managedDisk"`
	// netapp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#netapp AzurermProvider#netapp}
	Netapp interface{} `field:"optional" json:"netapp" yaml:"netapp"`
	// postgresql_flexible_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#postgresql_flexible_server AzurermProvider#postgresql_flexible_server}
	PostgresqlFlexibleServer interface{} `field:"optional" json:"postgresqlFlexibleServer" yaml:"postgresqlFlexibleServer"`
	// recovery_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#recovery_service AzurermProvider#recovery_service}
	RecoveryService interface{} `field:"optional" json:"recoveryService" yaml:"recoveryService"`
	// recovery_services_vaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#recovery_services_vaults AzurermProvider#recovery_services_vaults}
	RecoveryServicesVaults interface{} `field:"optional" json:"recoveryServicesVaults" yaml:"recoveryServicesVaults"`
	// resource_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#resource_group AzurermProvider#resource_group}
	ResourceGroup interface{} `field:"optional" json:"resourceGroup" yaml:"resourceGroup"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#storage AzurermProvider#storage}
	Storage interface{} `field:"optional" json:"storage" yaml:"storage"`
	// subscription block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#subscription AzurermProvider#subscription}
	Subscription interface{} `field:"optional" json:"subscription" yaml:"subscription"`
	// template_deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#template_deployment AzurermProvider#template_deployment}
	TemplateDeployment interface{} `field:"optional" json:"templateDeployment" yaml:"templateDeployment"`
	// virtual_machine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#virtual_machine AzurermProvider#virtual_machine}
	VirtualMachine interface{} `field:"optional" json:"virtualMachine" yaml:"virtualMachine"`
	// virtual_machine_scale_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs#virtual_machine_scale_set AzurermProvider#virtual_machine_scale_set}
	VirtualMachineScaleSet interface{} `field:"optional" json:"virtualMachineScaleSet" yaml:"virtualMachineScaleSet"`
}

