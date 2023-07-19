package datafactorytriggerschedule


type DataFactoryTriggerSchedulePipeline struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/data_factory_trigger_schedule#name DataFactoryTriggerSchedule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/data_factory_trigger_schedule#parameters DataFactoryTriggerSchedule#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

