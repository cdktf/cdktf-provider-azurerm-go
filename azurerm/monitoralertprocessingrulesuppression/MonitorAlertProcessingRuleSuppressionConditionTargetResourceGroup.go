package monitoralertprocessingrulesuppression


type MonitorAlertProcessingRuleSuppressionConditionTargetResourceGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/monitor_alert_processing_rule_suppression#operator MonitorAlertProcessingRuleSuppression#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/monitor_alert_processing_rule_suppression#values MonitorAlertProcessingRuleSuppression#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

