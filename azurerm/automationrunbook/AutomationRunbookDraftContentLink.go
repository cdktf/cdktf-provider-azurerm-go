package automationrunbook


type AutomationRunbookDraftContentLink struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#uri AutomationRunbook#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// hash block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#hash AutomationRunbook#hash}
	Hash *AutomationRunbookDraftContentLinkHash `field:"optional" json:"hash" yaml:"hash"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#version AutomationRunbook#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}
