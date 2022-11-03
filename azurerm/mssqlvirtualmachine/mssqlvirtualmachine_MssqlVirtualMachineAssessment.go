package mssqlvirtualmachine


type MssqlVirtualMachineAssessment struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#enabled MssqlVirtualMachine#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#run_immediately MssqlVirtualMachine#run_immediately}.
	RunImmediately interface{} `field:"optional" json:"runImmediately" yaml:"runImmediately"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#schedule MssqlVirtualMachine#schedule}
	Schedule *MssqlVirtualMachineAssessmentSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

