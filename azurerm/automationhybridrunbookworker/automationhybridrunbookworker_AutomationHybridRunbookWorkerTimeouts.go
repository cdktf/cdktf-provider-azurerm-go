package automationhybridrunbookworker


type AutomationHybridRunbookWorkerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_hybrid_runbook_worker#create AutomationHybridRunbookWorker#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_hybrid_runbook_worker#delete AutomationHybridRunbookWorker#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_hybrid_runbook_worker#read AutomationHybridRunbookWorker#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
