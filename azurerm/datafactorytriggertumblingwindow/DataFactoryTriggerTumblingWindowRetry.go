package datafactorytriggertumblingwindow


type DataFactoryTriggerTumblingWindowRetry struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_trigger_tumbling_window#count DataFactoryTriggerTumblingWindow#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_trigger_tumbling_window#interval DataFactoryTriggerTumblingWindow#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
}

