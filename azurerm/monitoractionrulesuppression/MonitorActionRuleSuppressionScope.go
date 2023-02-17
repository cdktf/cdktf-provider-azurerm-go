package monitoractionrulesuppression


type MonitorActionRuleSuppressionScope struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_rule_suppression#resource_ids MonitorActionRuleSuppression#resource_ids}.
	ResourceIds *[]*string `field:"required" json:"resourceIds" yaml:"resourceIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_rule_suppression#type MonitorActionRuleSuppression#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

