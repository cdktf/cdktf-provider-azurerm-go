package monitoractionrulesuppression


type MonitorActionRuleSuppressionSuppressionSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/monitor_action_rule_suppression#end_date_utc MonitorActionRuleSuppression#end_date_utc}.
	EndDateUtc *string `field:"required" json:"endDateUtc" yaml:"endDateUtc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/monitor_action_rule_suppression#start_date_utc MonitorActionRuleSuppression#start_date_utc}.
	StartDateUtc *string `field:"required" json:"startDateUtc" yaml:"startDateUtc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/monitor_action_rule_suppression#recurrence_monthly MonitorActionRuleSuppression#recurrence_monthly}.
	RecurrenceMonthly *[]*float64 `field:"optional" json:"recurrenceMonthly" yaml:"recurrenceMonthly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/monitor_action_rule_suppression#recurrence_weekly MonitorActionRuleSuppression#recurrence_weekly}.
	RecurrenceWeekly *[]*string `field:"optional" json:"recurrenceWeekly" yaml:"recurrenceWeekly"`
}

