package devtestschedule


type DevTestScheduleHourlyRecurrence struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_schedule#minute DevTestSchedule#minute}.
	Minute *float64 `field:"required" json:"minute" yaml:"minute"`
}

