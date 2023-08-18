package datafactorytriggerblobevent


type DataFactoryTriggerBlobEventPipeline struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/data_factory_trigger_blob_event#name DataFactoryTriggerBlobEvent#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/data_factory_trigger_blob_event#parameters DataFactoryTriggerBlobEvent#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

