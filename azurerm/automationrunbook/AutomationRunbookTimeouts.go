package automationrunbook


type AutomationRunbookTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#create AutomationRunbook#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#delete AutomationRunbook#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#read AutomationRunbook#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#update AutomationRunbook#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

