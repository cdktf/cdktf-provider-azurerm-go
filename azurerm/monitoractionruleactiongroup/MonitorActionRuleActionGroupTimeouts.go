package monitoractionruleactiongroup


type MonitorActionRuleActionGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_rule_action_group#create MonitorActionRuleActionGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_rule_action_group#delete MonitorActionRuleActionGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_rule_action_group#read MonitorActionRuleActionGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_rule_action_group#update MonitorActionRuleActionGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
