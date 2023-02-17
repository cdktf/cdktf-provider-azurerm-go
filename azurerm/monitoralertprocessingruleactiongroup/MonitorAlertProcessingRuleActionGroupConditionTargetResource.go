package monitoralertprocessingruleactiongroup


type MonitorAlertProcessingRuleActionGroupConditionTargetResource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#operator MonitorAlertProcessingRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_alert_processing_rule_action_group#values MonitorAlertProcessingRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

