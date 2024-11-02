// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlvirtualmachineavailabilitygrouplistener

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlVirtualMachineAvailabilityGroupListenerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/mssql_virtual_machine_availability_group_listener#name MssqlVirtualMachineAvailabilityGroupListener#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// replica block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/mssql_virtual_machine_availability_group_listener#replica MssqlVirtualMachineAvailabilityGroupListener#replica}
	Replica interface{} `field:"required" json:"replica" yaml:"replica"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/mssql_virtual_machine_availability_group_listener#sql_virtual_machine_group_id MssqlVirtualMachineAvailabilityGroupListener#sql_virtual_machine_group_id}.
	SqlVirtualMachineGroupId *string `field:"required" json:"sqlVirtualMachineGroupId" yaml:"sqlVirtualMachineGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/mssql_virtual_machine_availability_group_listener#availability_group_name MssqlVirtualMachineAvailabilityGroupListener#availability_group_name}.
	AvailabilityGroupName *string `field:"optional" json:"availabilityGroupName" yaml:"availabilityGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/mssql_virtual_machine_availability_group_listener#id MssqlVirtualMachineAvailabilityGroupListener#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// load_balancer_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/mssql_virtual_machine_availability_group_listener#load_balancer_configuration MssqlVirtualMachineAvailabilityGroupListener#load_balancer_configuration}
	LoadBalancerConfiguration *MssqlVirtualMachineAvailabilityGroupListenerLoadBalancerConfiguration `field:"optional" json:"loadBalancerConfiguration" yaml:"loadBalancerConfiguration"`
	// multi_subnet_ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/mssql_virtual_machine_availability_group_listener#multi_subnet_ip_configuration MssqlVirtualMachineAvailabilityGroupListener#multi_subnet_ip_configuration}
	MultiSubnetIpConfiguration interface{} `field:"optional" json:"multiSubnetIpConfiguration" yaml:"multiSubnetIpConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/mssql_virtual_machine_availability_group_listener#port MssqlVirtualMachineAvailabilityGroupListener#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/mssql_virtual_machine_availability_group_listener#timeouts MssqlVirtualMachineAvailabilityGroupListener#timeouts}
	Timeouts *MssqlVirtualMachineAvailabilityGroupListenerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

