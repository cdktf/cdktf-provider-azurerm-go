package datafactorytriggertumblingwindow


type DataFactoryTriggerTumblingWindowPipeline struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_trigger_tumbling_window#name DataFactoryTriggerTumblingWindow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_trigger_tumbling_window#parameters DataFactoryTriggerTumblingWindow#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}
