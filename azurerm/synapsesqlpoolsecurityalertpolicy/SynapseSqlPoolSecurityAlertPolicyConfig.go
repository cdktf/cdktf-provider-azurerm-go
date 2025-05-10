// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapsesqlpoolsecurityalertpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SynapseSqlPoolSecurityAlertPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/synapse_sql_pool_security_alert_policy#policy_state SynapseSqlPoolSecurityAlertPolicy#policy_state}.
	PolicyState *string `field:"required" json:"policyState" yaml:"policyState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/synapse_sql_pool_security_alert_policy#sql_pool_id SynapseSqlPoolSecurityAlertPolicy#sql_pool_id}.
	SqlPoolId *string `field:"required" json:"sqlPoolId" yaml:"sqlPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/synapse_sql_pool_security_alert_policy#disabled_alerts SynapseSqlPoolSecurityAlertPolicy#disabled_alerts}.
	DisabledAlerts *[]*string `field:"optional" json:"disabledAlerts" yaml:"disabledAlerts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/synapse_sql_pool_security_alert_policy#email_account_admins_enabled SynapseSqlPoolSecurityAlertPolicy#email_account_admins_enabled}.
	EmailAccountAdminsEnabled interface{} `field:"optional" json:"emailAccountAdminsEnabled" yaml:"emailAccountAdminsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/synapse_sql_pool_security_alert_policy#email_addresses SynapseSqlPoolSecurityAlertPolicy#email_addresses}.
	EmailAddresses *[]*string `field:"optional" json:"emailAddresses" yaml:"emailAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/synapse_sql_pool_security_alert_policy#id SynapseSqlPoolSecurityAlertPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/synapse_sql_pool_security_alert_policy#retention_days SynapseSqlPoolSecurityAlertPolicy#retention_days}.
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/synapse_sql_pool_security_alert_policy#storage_account_access_key SynapseSqlPoolSecurityAlertPolicy#storage_account_access_key}.
	StorageAccountAccessKey *string `field:"optional" json:"storageAccountAccessKey" yaml:"storageAccountAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/synapse_sql_pool_security_alert_policy#storage_endpoint SynapseSqlPoolSecurityAlertPolicy#storage_endpoint}.
	StorageEndpoint *string `field:"optional" json:"storageEndpoint" yaml:"storageEndpoint"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/synapse_sql_pool_security_alert_policy#timeouts SynapseSqlPoolSecurityAlertPolicy#timeouts}
	Timeouts *SynapseSqlPoolSecurityAlertPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

