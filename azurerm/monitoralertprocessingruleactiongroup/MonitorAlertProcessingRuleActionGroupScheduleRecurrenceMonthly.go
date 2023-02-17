package monitoralertprocessingruleactiongroup


type MonitorAlertProcessingRuleActionGroupScheduleRecurrenceMonthly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#days_of_month MonitorAlertProcessingRuleActionGroup#days_of_month}.
	DaysOfMonth *[]*float64 `field:"required" json:"daysOfMonth" yaml:"daysOfMonth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#end_time MonitorAlertProcessingRuleActionGroup#end_time}.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#start_time MonitorAlertProcessingRuleActionGroup#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

