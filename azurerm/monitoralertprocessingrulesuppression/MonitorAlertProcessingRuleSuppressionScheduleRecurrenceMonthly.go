package monitoralertprocessingrulesuppression


type MonitorAlertProcessingRuleSuppressionScheduleRecurrenceMonthly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_suppression#days_of_month MonitorAlertProcessingRuleSuppression#days_of_month}.
	DaysOfMonth *[]*float64 `field:"required" json:"daysOfMonth" yaml:"daysOfMonth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_suppression#end_time MonitorAlertProcessingRuleSuppression#end_time}.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_suppression#start_time MonitorAlertProcessingRuleSuppression#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}
