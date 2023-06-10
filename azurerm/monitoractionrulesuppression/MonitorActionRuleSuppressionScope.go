package monitoractionrulesuppression


type MonitorActionRuleSuppressionScope struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_rule_suppression#resource_ids MonitorActionRuleSuppression#resource_ids}.
	ResourceIds *[]*string `field:"required" json:"resourceIds" yaml:"resourceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/monitor_action_rule_suppression#type MonitorActionRuleSuppression#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

