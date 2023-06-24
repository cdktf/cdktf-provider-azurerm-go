package monitoractionruleactiongroup


type MonitorActionRuleActionGroupConditionMonitorService struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/monitor_action_rule_action_group#operator MonitorActionRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/monitor_action_rule_action_group#values MonitorActionRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

