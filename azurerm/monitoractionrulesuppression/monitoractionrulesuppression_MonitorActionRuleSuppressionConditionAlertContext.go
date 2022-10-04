package monitoractionrulesuppression


type MonitorActionRuleSuppressionConditionAlertContext struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_rule_suppression#operator MonitorActionRuleSuppression#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_rule_suppression#values MonitorActionRuleSuppression#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

