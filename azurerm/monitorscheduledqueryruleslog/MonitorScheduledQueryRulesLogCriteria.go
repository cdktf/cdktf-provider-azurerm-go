package monitorscheduledqueryruleslog


type MonitorScheduledQueryRulesLogCriteria struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_log#dimension MonitorScheduledQueryRulesLog#dimension}
	Dimension interface{} `field:"required" json:"dimension" yaml:"dimension"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_log#metric_name MonitorScheduledQueryRulesLog#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
}
