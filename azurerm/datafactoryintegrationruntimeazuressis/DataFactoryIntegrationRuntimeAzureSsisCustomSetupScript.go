package datafactoryintegrationruntimeazuressis


type DataFactoryIntegrationRuntimeAzureSsisCustomSetupScript struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/data_factory_integration_runtime_azure_ssis#blob_container_uri DataFactoryIntegrationRuntimeAzureSsis#blob_container_uri}.
	BlobContainerUri *string `field:"required" json:"blobContainerUri" yaml:"blobContainerUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/data_factory_integration_runtime_azure_ssis#sas_token DataFactoryIntegrationRuntimeAzureSsis#sas_token}.
	SasToken *string `field:"required" json:"sasToken" yaml:"sasToken"`
}

