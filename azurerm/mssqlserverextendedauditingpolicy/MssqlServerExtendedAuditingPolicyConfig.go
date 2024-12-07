// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlserverextendedauditingpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlServerExtendedAuditingPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mssql_server_extended_auditing_policy#server_id MssqlServerExtendedAuditingPolicy#server_id}.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mssql_server_extended_auditing_policy#audit_actions_and_groups MssqlServerExtendedAuditingPolicy#audit_actions_and_groups}.
	AuditActionsAndGroups *[]*string `field:"optional" json:"auditActionsAndGroups" yaml:"auditActionsAndGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mssql_server_extended_auditing_policy#enabled MssqlServerExtendedAuditingPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mssql_server_extended_auditing_policy#id MssqlServerExtendedAuditingPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mssql_server_extended_auditing_policy#log_monitoring_enabled MssqlServerExtendedAuditingPolicy#log_monitoring_enabled}.
	LogMonitoringEnabled interface{} `field:"optional" json:"logMonitoringEnabled" yaml:"logMonitoringEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mssql_server_extended_auditing_policy#predicate_expression MssqlServerExtendedAuditingPolicy#predicate_expression}.
	PredicateExpression *string `field:"optional" json:"predicateExpression" yaml:"predicateExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mssql_server_extended_auditing_policy#retention_in_days MssqlServerExtendedAuditingPolicy#retention_in_days}.
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mssql_server_extended_auditing_policy#storage_account_access_key MssqlServerExtendedAuditingPolicy#storage_account_access_key}.
	StorageAccountAccessKey *string `field:"optional" json:"storageAccountAccessKey" yaml:"storageAccountAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mssql_server_extended_auditing_policy#storage_account_access_key_is_secondary MssqlServerExtendedAuditingPolicy#storage_account_access_key_is_secondary}.
	StorageAccountAccessKeyIsSecondary interface{} `field:"optional" json:"storageAccountAccessKeyIsSecondary" yaml:"storageAccountAccessKeyIsSecondary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mssql_server_extended_auditing_policy#storage_account_subscription_id MssqlServerExtendedAuditingPolicy#storage_account_subscription_id}.
	StorageAccountSubscriptionId *string `field:"optional" json:"storageAccountSubscriptionId" yaml:"storageAccountSubscriptionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mssql_server_extended_auditing_policy#storage_endpoint MssqlServerExtendedAuditingPolicy#storage_endpoint}.
	StorageEndpoint *string `field:"optional" json:"storageEndpoint" yaml:"storageEndpoint"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mssql_server_extended_auditing_policy#timeouts MssqlServerExtendedAuditingPolicy#timeouts}
	Timeouts *MssqlServerExtendedAuditingPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

