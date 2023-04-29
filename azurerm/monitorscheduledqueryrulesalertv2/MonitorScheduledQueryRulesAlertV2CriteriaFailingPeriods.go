package monitorscheduledqueryrulesalertv2


type MonitorScheduledQueryRulesAlertV2CriteriaFailingPeriods struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/monitor_scheduled_query_rules_alert_v2#minimum_failing_periods_to_trigger_alert MonitorScheduledQueryRulesAlertV2#minimum_failing_periods_to_trigger_alert}.
	MinimumFailingPeriodsToTriggerAlert *float64 `field:"required" json:"minimumFailingPeriodsToTriggerAlert" yaml:"minimumFailingPeriodsToTriggerAlert"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/monitor_scheduled_query_rules_alert_v2#number_of_evaluation_periods MonitorScheduledQueryRulesAlertV2#number_of_evaluation_periods}.
	NumberOfEvaluationPeriods *float64 `field:"required" json:"numberOfEvaluationPeriods" yaml:"numberOfEvaluationPeriods"`
}

