package devtestschedule


type DevTestScheduleHourlyRecurrence struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/dev_test_schedule#minute DevTestSchedule#minute}.
	Minute *float64 `field:"required" json:"minute" yaml:"minute"`
}

