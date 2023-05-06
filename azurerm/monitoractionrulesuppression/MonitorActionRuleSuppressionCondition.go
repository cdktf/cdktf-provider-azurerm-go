package monitoractionrulesuppression


type MonitorActionRuleSuppressionCondition struct {
	// alert_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/monitor_action_rule_suppression#alert_context MonitorActionRuleSuppression#alert_context}
	AlertContext *MonitorActionRuleSuppressionConditionAlertContext `field:"optional" json:"alertContext" yaml:"alertContext"`
	// alert_rule_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/monitor_action_rule_suppression#alert_rule_id MonitorActionRuleSuppression#alert_rule_id}
	AlertRuleId *MonitorActionRuleSuppressionConditionAlertRuleId `field:"optional" json:"alertRuleId" yaml:"alertRuleId"`
	// description block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/monitor_action_rule_suppression#description MonitorActionRuleSuppression#description}
	Description *MonitorActionRuleSuppressionConditionDescription `field:"optional" json:"description" yaml:"description"`
	// monitor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/monitor_action_rule_suppression#monitor MonitorActionRuleSuppression#monitor}
	Monitor *MonitorActionRuleSuppressionConditionMonitor `field:"optional" json:"monitor" yaml:"monitor"`
	// monitor_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/monitor_action_rule_suppression#monitor_service MonitorActionRuleSuppression#monitor_service}
	MonitorService *MonitorActionRuleSuppressionConditionMonitorService `field:"optional" json:"monitorService" yaml:"monitorService"`
	// severity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/monitor_action_rule_suppression#severity MonitorActionRuleSuppression#severity}
	Severity *MonitorActionRuleSuppressionConditionSeverity `field:"optional" json:"severity" yaml:"severity"`
	// target_resource_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/monitor_action_rule_suppression#target_resource_type MonitorActionRuleSuppression#target_resource_type}
	TargetResourceType *MonitorActionRuleSuppressionConditionTargetResourceType `field:"optional" json:"targetResourceType" yaml:"targetResourceType"`
}

