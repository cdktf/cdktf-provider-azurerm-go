// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactoryintegrationruntimeazuressis


type DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupComponentKeyVaultLicense struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/data_factory_integration_runtime_azure_ssis#linked_service_name DataFactoryIntegrationRuntimeAzureSsis#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/data_factory_integration_runtime_azure_ssis#secret_name DataFactoryIntegrationRuntimeAzureSsis#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/data_factory_integration_runtime_azure_ssis#parameters DataFactoryIntegrationRuntimeAzureSsis#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/data_factory_integration_runtime_azure_ssis#secret_version DataFactoryIntegrationRuntimeAzureSsis#secret_version}.
	SecretVersion *string `field:"optional" json:"secretVersion" yaml:"secretVersion"`
}

