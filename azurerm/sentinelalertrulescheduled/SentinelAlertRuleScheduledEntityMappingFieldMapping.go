// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulescheduled


type SentinelAlertRuleScheduledEntityMappingFieldMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/sentinel_alert_rule_scheduled#column_name SentinelAlertRuleScheduled#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/sentinel_alert_rule_scheduled#identifier SentinelAlertRuleScheduled#identifier}.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

