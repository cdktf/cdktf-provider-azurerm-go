package datadogmonitortagrule


type DatadogMonitorTagRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor_tag_rule#create DatadogMonitorTagRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor_tag_rule#delete DatadogMonitorTagRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor_tag_rule#read DatadogMonitorTagRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor_tag_rule#update DatadogMonitorTagRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
