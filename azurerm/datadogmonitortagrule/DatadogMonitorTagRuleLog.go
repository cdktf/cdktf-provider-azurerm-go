package datadogmonitortagrule


type DatadogMonitorTagRuleLog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/datadog_monitor_tag_rule#aad_log_enabled DatadogMonitorTagRule#aad_log_enabled}.
	AadLogEnabled interface{} `field:"optional" json:"aadLogEnabled" yaml:"aadLogEnabled"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/datadog_monitor_tag_rule#filter DatadogMonitorTagRule#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/datadog_monitor_tag_rule#resource_log_enabled DatadogMonitorTagRule#resource_log_enabled}.
	ResourceLogEnabled interface{} `field:"optional" json:"resourceLogEnabled" yaml:"resourceLogEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/datadog_monitor_tag_rule#subscription_log_enabled DatadogMonitorTagRule#subscription_log_enabled}.
	SubscriptionLogEnabled interface{} `field:"optional" json:"subscriptionLogEnabled" yaml:"subscriptionLogEnabled"`
}

