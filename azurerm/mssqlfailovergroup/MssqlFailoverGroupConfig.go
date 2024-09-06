// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlfailovergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlFailoverGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/mssql_failover_group#name MssqlFailoverGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// partner_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/mssql_failover_group#partner_server MssqlFailoverGroup#partner_server}
	PartnerServer interface{} `field:"required" json:"partnerServer" yaml:"partnerServer"`
	// read_write_endpoint_failover_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/mssql_failover_group#read_write_endpoint_failover_policy MssqlFailoverGroup#read_write_endpoint_failover_policy}
	ReadWriteEndpointFailoverPolicy *MssqlFailoverGroupReadWriteEndpointFailoverPolicy `field:"required" json:"readWriteEndpointFailoverPolicy" yaml:"readWriteEndpointFailoverPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/mssql_failover_group#server_id MssqlFailoverGroup#server_id}.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/mssql_failover_group#databases MssqlFailoverGroup#databases}.
	Databases *[]*string `field:"optional" json:"databases" yaml:"databases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/mssql_failover_group#id MssqlFailoverGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/mssql_failover_group#readonly_endpoint_failover_policy_enabled MssqlFailoverGroup#readonly_endpoint_failover_policy_enabled}.
	ReadonlyEndpointFailoverPolicyEnabled interface{} `field:"optional" json:"readonlyEndpointFailoverPolicyEnabled" yaml:"readonlyEndpointFailoverPolicyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/mssql_failover_group#tags MssqlFailoverGroup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/mssql_failover_group#timeouts MssqlFailoverGroup#timeouts}
	Timeouts *MssqlFailoverGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

