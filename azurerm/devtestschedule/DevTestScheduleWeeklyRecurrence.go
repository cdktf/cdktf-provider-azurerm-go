package devtestschedule


type DevTestScheduleWeeklyRecurrence struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/dev_test_schedule#time DevTestSchedule#time}.
	Time *string `field:"required" json:"time" yaml:"time"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/dev_test_schedule#week_days DevTestSchedule#week_days}.
	WeekDays *[]*string `field:"optional" json:"weekDays" yaml:"weekDays"`
}

