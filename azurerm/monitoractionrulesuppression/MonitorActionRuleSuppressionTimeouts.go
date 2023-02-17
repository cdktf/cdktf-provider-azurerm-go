package monitoractionrulesuppression


type MonitorActionRuleSuppressionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_rule_suppression#create MonitorActionRuleSuppression#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_rule_suppression#delete MonitorActionRuleSuppression#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_rule_suppression#read MonitorActionRuleSuppression#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_rule_suppression#update MonitorActionRuleSuppression#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

