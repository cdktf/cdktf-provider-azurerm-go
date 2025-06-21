// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapsesqlpoolworkloadgroup


type SynapseSqlPoolWorkloadGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/synapse_sql_pool_workload_group#create SynapseSqlPoolWorkloadGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/synapse_sql_pool_workload_group#delete SynapseSqlPoolWorkloadGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/synapse_sql_pool_workload_group#read SynapseSqlPoolWorkloadGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/synapse_sql_pool_workload_group#update SynapseSqlPoolWorkloadGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

