package datafactoryintegrationruntimemanaged


type DataFactoryIntegrationRuntimeManagedCustomSetupScript struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_integration_runtime_managed#blob_container_uri DataFactoryIntegrationRuntimeManaged#blob_container_uri}.
	BlobContainerUri *string `field:"required" json:"blobContainerUri" yaml:"blobContainerUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_integration_runtime_managed#sas_token DataFactoryIntegrationRuntimeManaged#sas_token}.
	SasToken *string `field:"required" json:"sasToken" yaml:"sasToken"`
}
