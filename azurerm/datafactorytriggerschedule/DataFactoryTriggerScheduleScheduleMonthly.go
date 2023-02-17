package datafactorytriggerschedule


type DataFactoryTriggerScheduleScheduleMonthly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_trigger_schedule#weekday DataFactoryTriggerSchedule#weekday}.
	Weekday *string `field:"required" json:"weekday" yaml:"weekday"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_trigger_schedule#week DataFactoryTriggerSchedule#week}.
	Week *float64 `field:"optional" json:"week" yaml:"week"`
}

