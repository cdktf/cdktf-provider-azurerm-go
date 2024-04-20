// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelalertrulescheduled


type SentinelAlertRuleScheduledEntityMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/sentinel_alert_rule_scheduled#entity_type SentinelAlertRuleScheduled#entity_type}.
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// field_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/sentinel_alert_rule_scheduled#field_mapping SentinelAlertRuleScheduled#field_mapping}
	FieldMapping interface{} `field:"required" json:"fieldMapping" yaml:"fieldMapping"`
}

