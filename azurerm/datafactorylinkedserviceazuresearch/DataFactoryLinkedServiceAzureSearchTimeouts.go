package datafactorylinkedserviceazuresearch


type DataFactoryLinkedServiceAzureSearchTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_search#create DataFactoryLinkedServiceAzureSearch#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_search#delete DataFactoryLinkedServiceAzureSearch#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_search#read DataFactoryLinkedServiceAzureSearch#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_service_azure_search#update DataFactoryLinkedServiceAzureSearch#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
