package monitoralertprocessingruleactiongroup


type MonitorAlertProcessingRuleActionGroupScheduleRecurrenceDaily struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#end_time MonitorAlertProcessingRuleActionGroup#end_time}.
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#start_time MonitorAlertProcessingRuleActionGroup#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}
