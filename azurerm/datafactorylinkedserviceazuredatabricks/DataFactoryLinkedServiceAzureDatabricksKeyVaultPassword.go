package datafactorylinkedserviceazuredatabricks


type DataFactoryLinkedServiceAzureDatabricksKeyVaultPassword struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_linked_service_azure_databricks#linked_service_name DataFactoryLinkedServiceAzureDatabricks#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_linked_service_azure_databricks#secret_name DataFactoryLinkedServiceAzureDatabricks#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

