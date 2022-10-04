package automationrunbook


type AutomationRunbookJobSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#job_schedule_id AutomationRunbook#job_schedule_id}.
	JobScheduleId *string `field:"optional" json:"jobScheduleId" yaml:"jobScheduleId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#parameters AutomationRunbook#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#run_on AutomationRunbook#run_on}.
	RunOn *string `field:"optional" json:"runOn" yaml:"runOn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#schedule_name AutomationRunbook#schedule_name}.
	ScheduleName *string `field:"optional" json:"scheduleName" yaml:"scheduleName"`
}

