// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlmanagedinstancefailovergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlManagedInstanceFailoverGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/mssql_managed_instance_failover_group#location MssqlManagedInstanceFailoverGroup#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/mssql_managed_instance_failover_group#managed_instance_id MssqlManagedInstanceFailoverGroup#managed_instance_id}.
	ManagedInstanceId *string `field:"required" json:"managedInstanceId" yaml:"managedInstanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/mssql_managed_instance_failover_group#name MssqlManagedInstanceFailoverGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/mssql_managed_instance_failover_group#partner_managed_instance_id MssqlManagedInstanceFailoverGroup#partner_managed_instance_id}.
	PartnerManagedInstanceId *string `field:"required" json:"partnerManagedInstanceId" yaml:"partnerManagedInstanceId"`
	// read_write_endpoint_failover_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/mssql_managed_instance_failover_group#read_write_endpoint_failover_policy MssqlManagedInstanceFailoverGroup#read_write_endpoint_failover_policy}
	ReadWriteEndpointFailoverPolicy *MssqlManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicy `field:"required" json:"readWriteEndpointFailoverPolicy" yaml:"readWriteEndpointFailoverPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/mssql_managed_instance_failover_group#id MssqlManagedInstanceFailoverGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/mssql_managed_instance_failover_group#readonly_endpoint_failover_policy_enabled MssqlManagedInstanceFailoverGroup#readonly_endpoint_failover_policy_enabled}.
	ReadonlyEndpointFailoverPolicyEnabled interface{} `field:"optional" json:"readonlyEndpointFailoverPolicyEnabled" yaml:"readonlyEndpointFailoverPolicyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/mssql_managed_instance_failover_group#secondary_type MssqlManagedInstanceFailoverGroup#secondary_type}.
	SecondaryType *string `field:"optional" json:"secondaryType" yaml:"secondaryType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/mssql_managed_instance_failover_group#timeouts MssqlManagedInstanceFailoverGroup#timeouts}
	Timeouts *MssqlManagedInstanceFailoverGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

