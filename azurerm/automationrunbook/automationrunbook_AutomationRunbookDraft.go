package automationrunbook


type AutomationRunbookDraft struct {
	// content_link block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#content_link AutomationRunbook#content_link}
	ContentLink *AutomationRunbookDraftContentLink `field:"optional" json:"contentLink" yaml:"contentLink"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#edit_mode_enabled AutomationRunbook#edit_mode_enabled}.
	EditModeEnabled interface{} `field:"optional" json:"editModeEnabled" yaml:"editModeEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#output_types AutomationRunbook#output_types}.
	OutputTypes *[]*string `field:"optional" json:"outputTypes" yaml:"outputTypes"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_runbook#parameters AutomationRunbook#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}
