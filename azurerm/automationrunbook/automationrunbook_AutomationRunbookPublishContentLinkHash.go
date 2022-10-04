package automationrunbook


type AutomationRunbookPublishContentLinkHash struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#algorithm AutomationRunbook#algorithm}.
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#value AutomationRunbook#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

