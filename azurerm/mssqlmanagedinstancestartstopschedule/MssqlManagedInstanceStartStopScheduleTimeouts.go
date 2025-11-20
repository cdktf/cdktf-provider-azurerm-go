// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlmanagedinstancestartstopschedule


type MssqlManagedInstanceStartStopScheduleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/mssql_managed_instance_start_stop_schedule#create MssqlManagedInstanceStartStopSchedule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/mssql_managed_instance_start_stop_schedule#delete MssqlManagedInstanceStartStopSchedule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/mssql_managed_instance_start_stop_schedule#read MssqlManagedInstanceStartStopSchedule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/mssql_managed_instance_start_stop_schedule#update MssqlManagedInstanceStartStopSchedule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

