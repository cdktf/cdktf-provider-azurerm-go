package monitoralertprocessingrulesuppression


type MonitorAlertProcessingRuleSuppressionScheduleRecurrence struct {
	// daily block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_suppression#daily MonitorAlertProcessingRuleSuppression#daily}
	Daily interface{} `field:"optional" json:"daily" yaml:"daily"`
	// monthly block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_suppression#monthly MonitorAlertProcessingRuleSuppression#monthly}
	Monthly interface{} `field:"optional" json:"monthly" yaml:"monthly"`
	// weekly block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_suppression#weekly MonitorAlertProcessingRuleSuppression#weekly}
	Weekly interface{} `field:"optional" json:"weekly" yaml:"weekly"`
}

