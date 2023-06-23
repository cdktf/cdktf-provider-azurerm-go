package mssqlvirtualmachine


type MssqlVirtualMachineAssessmentSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/mssql_virtual_machine#day_of_week MssqlVirtualMachine#day_of_week}.
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/mssql_virtual_machine#start_time MssqlVirtualMachine#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/mssql_virtual_machine#monthly_occurrence MssqlVirtualMachine#monthly_occurrence}.
	MonthlyOccurrence *float64 `field:"optional" json:"monthlyOccurrence" yaml:"monthlyOccurrence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/mssql_virtual_machine#weekly_interval MssqlVirtualMachine#weekly_interval}.
	WeeklyInterval *float64 `field:"optional" json:"weeklyInterval" yaml:"weeklyInterval"`
}

