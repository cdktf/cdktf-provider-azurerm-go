package automationschedule


type AutomationScheduleMonthlyOccurrence struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/automation_schedule#day AutomationSchedule#day}.
	Day *string `field:"required" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/automation_schedule#occurrence AutomationSchedule#occurrence}.
	Occurrence *float64 `field:"required" json:"occurrence" yaml:"occurrence"`
}

