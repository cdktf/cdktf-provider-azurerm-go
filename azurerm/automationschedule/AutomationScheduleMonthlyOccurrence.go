package automationschedule


type AutomationScheduleMonthlyOccurrence struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_schedule#day AutomationSchedule#day}.
	Day *string `field:"required" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_schedule#occurrence AutomationSchedule#occurrence}.
	Occurrence *float64 `field:"required" json:"occurrence" yaml:"occurrence"`
}

