// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlvirtualmachineavailabilitygrouplistener


type MssqlVirtualMachineAvailabilityGroupListenerReplica struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/mssql_virtual_machine_availability_group_listener#commit MssqlVirtualMachineAvailabilityGroupListener#commit}.
	Commit *string `field:"required" json:"commit" yaml:"commit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/mssql_virtual_machine_availability_group_listener#failover_mode MssqlVirtualMachineAvailabilityGroupListener#failover_mode}.
	FailoverMode *string `field:"required" json:"failoverMode" yaml:"failoverMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/mssql_virtual_machine_availability_group_listener#readable_secondary MssqlVirtualMachineAvailabilityGroupListener#readable_secondary}.
	ReadableSecondary *string `field:"required" json:"readableSecondary" yaml:"readableSecondary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/mssql_virtual_machine_availability_group_listener#role MssqlVirtualMachineAvailabilityGroupListener#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/mssql_virtual_machine_availability_group_listener#sql_virtual_machine_id MssqlVirtualMachineAvailabilityGroupListener#sql_virtual_machine_id}.
	SqlVirtualMachineId *string `field:"required" json:"sqlVirtualMachineId" yaml:"sqlVirtualMachineId"`
}

