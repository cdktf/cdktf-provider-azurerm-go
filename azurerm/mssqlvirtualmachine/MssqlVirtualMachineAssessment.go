// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlvirtualmachine


type MssqlVirtualMachineAssessment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/mssql_virtual_machine#enabled MssqlVirtualMachine#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/mssql_virtual_machine#run_immediately MssqlVirtualMachine#run_immediately}.
	RunImmediately interface{} `field:"optional" json:"runImmediately" yaml:"runImmediately"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/mssql_virtual_machine#schedule MssqlVirtualMachine#schedule}
	Schedule *MssqlVirtualMachineAssessmentSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

