package datafactorytriggercustomevent


type DataFactoryTriggerCustomEventPipeline struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_trigger_custom_event#name DataFactoryTriggerCustomEvent#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_trigger_custom_event#parameters DataFactoryTriggerCustomEvent#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}
