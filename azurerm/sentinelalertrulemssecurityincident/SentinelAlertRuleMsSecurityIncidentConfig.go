// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulemssecurityincident

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelAlertRuleMsSecurityIncidentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/sentinel_alert_rule_ms_security_incident#display_name SentinelAlertRuleMsSecurityIncident#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/sentinel_alert_rule_ms_security_incident#log_analytics_workspace_id SentinelAlertRuleMsSecurityIncident#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/sentinel_alert_rule_ms_security_incident#name SentinelAlertRuleMsSecurityIncident#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/sentinel_alert_rule_ms_security_incident#product_filter SentinelAlertRuleMsSecurityIncident#product_filter}.
	ProductFilter *string `field:"required" json:"productFilter" yaml:"productFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/sentinel_alert_rule_ms_security_incident#severity_filter SentinelAlertRuleMsSecurityIncident#severity_filter}.
	SeverityFilter *[]*string `field:"required" json:"severityFilter" yaml:"severityFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/sentinel_alert_rule_ms_security_incident#alert_rule_template_guid SentinelAlertRuleMsSecurityIncident#alert_rule_template_guid}.
	AlertRuleTemplateGuid *string `field:"optional" json:"alertRuleTemplateGuid" yaml:"alertRuleTemplateGuid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/sentinel_alert_rule_ms_security_incident#description SentinelAlertRuleMsSecurityIncident#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/sentinel_alert_rule_ms_security_incident#display_name_exclude_filter SentinelAlertRuleMsSecurityIncident#display_name_exclude_filter}.
	DisplayNameExcludeFilter *[]*string `field:"optional" json:"displayNameExcludeFilter" yaml:"displayNameExcludeFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/sentinel_alert_rule_ms_security_incident#display_name_filter SentinelAlertRuleMsSecurityIncident#display_name_filter}.
	DisplayNameFilter *[]*string `field:"optional" json:"displayNameFilter" yaml:"displayNameFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/sentinel_alert_rule_ms_security_incident#enabled SentinelAlertRuleMsSecurityIncident#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/sentinel_alert_rule_ms_security_incident#id SentinelAlertRuleMsSecurityIncident#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/sentinel_alert_rule_ms_security_incident#timeouts SentinelAlertRuleMsSecurityIncident#timeouts}
	Timeouts *SentinelAlertRuleMsSecurityIncidentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

