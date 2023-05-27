package logicappworkflow


type LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicy struct {
	// claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/logic_app_workflow#claim LogicAppWorkflow#claim}
	Claim interface{} `field:"required" json:"claim" yaml:"claim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/logic_app_workflow#name LogicAppWorkflow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

