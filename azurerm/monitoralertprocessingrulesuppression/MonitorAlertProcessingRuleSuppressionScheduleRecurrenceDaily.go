package monitoralertprocessingrulesuppression


type MonitorAlertProcessingRuleSuppressionScheduleRecurrenceDaily struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_suppression#end_time MonitorAlertProcessingRuleSuppression#end_time}.
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_suppression#start_time MonitorAlertProcessingRuleSuppression#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

