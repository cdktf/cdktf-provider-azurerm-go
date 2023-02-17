package sentinelalertrulenrt


type SentinelAlertRuleNrtEntityMapping struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_nrt#entity_type SentinelAlertRuleNrt#entity_type}.
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// field_mapping block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_nrt#field_mapping SentinelAlertRuleNrt#field_mapping}
	FieldMapping interface{} `field:"required" json:"fieldMapping" yaml:"fieldMapping"`
}

