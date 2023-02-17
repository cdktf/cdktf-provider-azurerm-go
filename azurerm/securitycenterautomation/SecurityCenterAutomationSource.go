package securitycenterautomation


type SecurityCenterAutomationSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_automation#event_source SecurityCenterAutomation#event_source}.
	EventSource *string `field:"required" json:"eventSource" yaml:"eventSource"`
	// rule_set block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_automation#rule_set SecurityCenterAutomation#rule_set}
	RuleSet interface{} `field:"optional" json:"ruleSet" yaml:"ruleSet"`
}

