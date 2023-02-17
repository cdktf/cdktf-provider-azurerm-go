package devtestschedule


type DevTestScheduleDailyRecurrence struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_schedule#time DevTestSchedule#time}.
	Time *string `field:"required" json:"time" yaml:"time"`
}

