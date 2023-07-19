package datafactorylinkedservicesnowflake


type DataFactoryLinkedServiceSnowflakeKeyVaultPassword struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/data_factory_linked_service_snowflake#linked_service_name DataFactoryLinkedServiceSnowflake#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/data_factory_linked_service_snowflake#secret_name DataFactoryLinkedServiceSnowflake#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

