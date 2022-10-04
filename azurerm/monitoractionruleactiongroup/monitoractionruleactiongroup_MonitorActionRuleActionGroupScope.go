package monitoractionruleactiongroup


type MonitorActionRuleActionGroupScope struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_rule_action_group#resource_ids MonitorActionRuleActionGroup#resource_ids}.
	ResourceIds *[]*string `field:"required" json:"resourceIds" yaml:"resourceIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_rule_action_group#type MonitorActionRuleActionGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

