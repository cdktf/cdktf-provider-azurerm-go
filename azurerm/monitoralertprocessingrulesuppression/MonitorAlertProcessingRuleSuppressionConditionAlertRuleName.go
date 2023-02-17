package monitoralertprocessingrulesuppression


type MonitorAlertProcessingRuleSuppressionConditionAlertRuleName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_suppression#operator MonitorAlertProcessingRuleSuppression#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_suppression#values MonitorAlertProcessingRuleSuppression#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

