package sentinelalertrulescheduled


type SentinelAlertRuleScheduledAlertDetailsOverrideDynamicProperty struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#name SentinelAlertRuleScheduled#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_scheduled#value SentinelAlertRuleScheduled#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

