package datafactoryintegrationruntimeazure


type DataFactoryIntegrationRuntimeAzureTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_integration_runtime_azure#create DataFactoryIntegrationRuntimeAzure#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_integration_runtime_azure#delete DataFactoryIntegrationRuntimeAzure#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_integration_runtime_azure#read DataFactoryIntegrationRuntimeAzure#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_integration_runtime_azure#update DataFactoryIntegrationRuntimeAzure#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
