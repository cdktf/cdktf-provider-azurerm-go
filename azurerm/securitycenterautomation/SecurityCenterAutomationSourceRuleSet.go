package securitycenterautomation


type SecurityCenterAutomationSourceRuleSet struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_automation#rule SecurityCenterAutomation#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
}

