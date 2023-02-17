package logicappworkflow


type LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaim struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_workflow#name LogicAppWorkflow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_workflow#value LogicAppWorkflow#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

