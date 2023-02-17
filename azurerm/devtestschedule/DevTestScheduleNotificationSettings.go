package devtestschedule


type DevTestScheduleNotificationSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_schedule#status DevTestSchedule#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_schedule#time_in_minutes DevTestSchedule#time_in_minutes}.
	TimeInMinutes *float64 `field:"optional" json:"timeInMinutes" yaml:"timeInMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_schedule#webhook_url DevTestSchedule#webhook_url}.
	WebhookUrl *string `field:"optional" json:"webhookUrl" yaml:"webhookUrl"`
}

