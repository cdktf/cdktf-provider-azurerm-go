package datafactorylinkedcustomservice


type DataFactoryLinkedCustomServiceIntegrationRuntime struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_custom_service#name DataFactoryLinkedCustomService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_linked_custom_service#parameters DataFactoryLinkedCustomService#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}
