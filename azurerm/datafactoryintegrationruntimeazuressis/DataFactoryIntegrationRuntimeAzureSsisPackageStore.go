package datafactoryintegrationruntimeazuressis


type DataFactoryIntegrationRuntimeAzureSsisPackageStore struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_integration_runtime_azure_ssis#linked_service_name DataFactoryIntegrationRuntimeAzureSsis#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_integration_runtime_azure_ssis#name DataFactoryIntegrationRuntimeAzureSsis#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}
