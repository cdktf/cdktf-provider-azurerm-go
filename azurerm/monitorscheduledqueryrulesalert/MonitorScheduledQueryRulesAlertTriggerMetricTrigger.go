package monitorscheduledqueryrulesalert


type MonitorScheduledQueryRulesAlertTriggerMetricTrigger struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#metric_column MonitorScheduledQueryRulesAlert#metric_column}.
	MetricColumn *string `field:"required" json:"metricColumn" yaml:"metricColumn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#metric_trigger_type MonitorScheduledQueryRulesAlert#metric_trigger_type}.
	MetricTriggerType *string `field:"required" json:"metricTriggerType" yaml:"metricTriggerType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#operator MonitorScheduledQueryRulesAlert#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_alert#threshold MonitorScheduledQueryRulesAlert#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
}

