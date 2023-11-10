// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlmanagedinstancefailovergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SqlManagedInstanceFailoverGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/sql_managed_instance_failover_group#location SqlManagedInstanceFailoverGroup#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/sql_managed_instance_failover_group#managed_instance_name SqlManagedInstanceFailoverGroup#managed_instance_name}.
	ManagedInstanceName *string `field:"required" json:"managedInstanceName" yaml:"managedInstanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/sql_managed_instance_failover_group#name SqlManagedInstanceFailoverGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/sql_managed_instance_failover_group#partner_managed_instance_id SqlManagedInstanceFailoverGroup#partner_managed_instance_id}.
	PartnerManagedInstanceId *string `field:"required" json:"partnerManagedInstanceId" yaml:"partnerManagedInstanceId"`
	// read_write_endpoint_failover_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/sql_managed_instance_failover_group#read_write_endpoint_failover_policy SqlManagedInstanceFailoverGroup#read_write_endpoint_failover_policy}
	ReadWriteEndpointFailoverPolicy *SqlManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicy `field:"required" json:"readWriteEndpointFailoverPolicy" yaml:"readWriteEndpointFailoverPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/sql_managed_instance_failover_group#resource_group_name SqlManagedInstanceFailoverGroup#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/sql_managed_instance_failover_group#id SqlManagedInstanceFailoverGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/sql_managed_instance_failover_group#readonly_endpoint_failover_policy_enabled SqlManagedInstanceFailoverGroup#readonly_endpoint_failover_policy_enabled}.
	ReadonlyEndpointFailoverPolicyEnabled interface{} `field:"optional" json:"readonlyEndpointFailoverPolicyEnabled" yaml:"readonlyEndpointFailoverPolicyEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/sql_managed_instance_failover_group#timeouts SqlManagedInstanceFailoverGroup#timeouts}
	Timeouts *SqlManagedInstanceFailoverGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

