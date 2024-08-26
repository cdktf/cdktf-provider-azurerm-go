// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactoryintegrationruntimeazuressis


type DataFactoryIntegrationRuntimeAzureSsisProxy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory_integration_runtime_azure_ssis#self_hosted_integration_runtime_name DataFactoryIntegrationRuntimeAzureSsis#self_hosted_integration_runtime_name}.
	SelfHostedIntegrationRuntimeName *string `field:"required" json:"selfHostedIntegrationRuntimeName" yaml:"selfHostedIntegrationRuntimeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory_integration_runtime_azure_ssis#staging_storage_linked_service_name DataFactoryIntegrationRuntimeAzureSsis#staging_storage_linked_service_name}.
	StagingStorageLinkedServiceName *string `field:"required" json:"stagingStorageLinkedServiceName" yaml:"stagingStorageLinkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory_integration_runtime_azure_ssis#path DataFactoryIntegrationRuntimeAzureSsis#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

