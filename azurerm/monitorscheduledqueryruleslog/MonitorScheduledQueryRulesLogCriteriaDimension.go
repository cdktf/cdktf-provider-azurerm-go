package monitorscheduledqueryruleslog


type MonitorScheduledQueryRulesLogCriteriaDimension struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_log#name MonitorScheduledQueryRulesLog#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_log#values MonitorScheduledQueryRulesLog#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_scheduled_query_rules_log#operator MonitorScheduledQueryRulesLog#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}
