// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapsesqlpoolextendedauditingpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SynapseSqlPoolExtendedAuditingPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/synapse_sql_pool_extended_auditing_policy#sql_pool_id SynapseSqlPoolExtendedAuditingPolicy#sql_pool_id}.
	SqlPoolId *string `field:"required" json:"sqlPoolId" yaml:"sqlPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/synapse_sql_pool_extended_auditing_policy#id SynapseSqlPoolExtendedAuditingPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/synapse_sql_pool_extended_auditing_policy#log_monitoring_enabled SynapseSqlPoolExtendedAuditingPolicy#log_monitoring_enabled}.
	LogMonitoringEnabled interface{} `field:"optional" json:"logMonitoringEnabled" yaml:"logMonitoringEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/synapse_sql_pool_extended_auditing_policy#retention_in_days SynapseSqlPoolExtendedAuditingPolicy#retention_in_days}.
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/synapse_sql_pool_extended_auditing_policy#storage_account_access_key SynapseSqlPoolExtendedAuditingPolicy#storage_account_access_key}.
	StorageAccountAccessKey *string `field:"optional" json:"storageAccountAccessKey" yaml:"storageAccountAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/synapse_sql_pool_extended_auditing_policy#storage_account_access_key_is_secondary SynapseSqlPoolExtendedAuditingPolicy#storage_account_access_key_is_secondary}.
	StorageAccountAccessKeyIsSecondary interface{} `field:"optional" json:"storageAccountAccessKeyIsSecondary" yaml:"storageAccountAccessKeyIsSecondary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/synapse_sql_pool_extended_auditing_policy#storage_endpoint SynapseSqlPoolExtendedAuditingPolicy#storage_endpoint}.
	StorageEndpoint *string `field:"optional" json:"storageEndpoint" yaml:"storageEndpoint"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/synapse_sql_pool_extended_auditing_policy#timeouts SynapseSqlPoolExtendedAuditingPolicy#timeouts}
	Timeouts *SynapseSqlPoolExtendedAuditingPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

