// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactoryintegrationruntimeazuressis


type DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetup struct {
	// command_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/data_factory_integration_runtime_azure_ssis#command_key DataFactoryIntegrationRuntimeAzureSsis#command_key}
	CommandKey interface{} `field:"optional" json:"commandKey" yaml:"commandKey"`
	// component block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/data_factory_integration_runtime_azure_ssis#component DataFactoryIntegrationRuntimeAzureSsis#component}
	Component interface{} `field:"optional" json:"component" yaml:"component"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/data_factory_integration_runtime_azure_ssis#environment DataFactoryIntegrationRuntimeAzureSsis#environment}.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/data_factory_integration_runtime_azure_ssis#powershell_version DataFactoryIntegrationRuntimeAzureSsis#powershell_version}.
	PowershellVersion *string `field:"optional" json:"powershellVersion" yaml:"powershellVersion"`
}

