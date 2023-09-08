// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlvirtualmachineavailabilitygrouplistener


type MssqlVirtualMachineAvailabilityGroupListenerLoadBalancerConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mssql_virtual_machine_availability_group_listener#load_balancer_id MssqlVirtualMachineAvailabilityGroupListener#load_balancer_id}.
	LoadBalancerId *string `field:"required" json:"loadBalancerId" yaml:"loadBalancerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mssql_virtual_machine_availability_group_listener#private_ip_address MssqlVirtualMachineAvailabilityGroupListener#private_ip_address}.
	PrivateIpAddress *string `field:"required" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mssql_virtual_machine_availability_group_listener#probe_port MssqlVirtualMachineAvailabilityGroupListener#probe_port}.
	ProbePort *float64 `field:"required" json:"probePort" yaml:"probePort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mssql_virtual_machine_availability_group_listener#sql_virtual_machine_ids MssqlVirtualMachineAvailabilityGroupListener#sql_virtual_machine_ids}.
	SqlVirtualMachineIds *[]*string `field:"required" json:"sqlVirtualMachineIds" yaml:"sqlVirtualMachineIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/mssql_virtual_machine_availability_group_listener#subnet_id MssqlVirtualMachineAvailabilityGroupListener#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

