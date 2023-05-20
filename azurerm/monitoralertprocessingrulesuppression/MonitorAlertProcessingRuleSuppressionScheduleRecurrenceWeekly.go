package monitoralertprocessingrulesuppression


type MonitorAlertProcessingRuleSuppressionScheduleRecurrenceWeekly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/monitor_alert_processing_rule_suppression#days_of_week MonitorAlertProcessingRuleSuppression#days_of_week}.
	DaysOfWeek *[]*string `field:"required" json:"daysOfWeek" yaml:"daysOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/monitor_alert_processing_rule_suppression#end_time MonitorAlertProcessingRuleSuppression#end_time}.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/monitor_alert_processing_rule_suppression#start_time MonitorAlertProcessingRuleSuppression#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

