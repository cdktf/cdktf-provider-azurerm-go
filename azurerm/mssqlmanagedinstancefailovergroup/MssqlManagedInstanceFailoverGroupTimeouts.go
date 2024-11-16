// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlmanagedinstancefailovergroup


type MssqlManagedInstanceFailoverGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/mssql_managed_instance_failover_group#create MssqlManagedInstanceFailoverGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/mssql_managed_instance_failover_group#delete MssqlManagedInstanceFailoverGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/mssql_managed_instance_failover_group#read MssqlManagedInstanceFailoverGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/mssql_managed_instance_failover_group#update MssqlManagedInstanceFailoverGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

