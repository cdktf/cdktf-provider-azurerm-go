package logicappactionhttp


type LogicAppActionHttpRunAfter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/logic_app_action_http#action_name LogicAppActionHttp#action_name}.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/logic_app_action_http#action_result LogicAppActionHttp#action_result}.
	ActionResult *string `field:"required" json:"actionResult" yaml:"actionResult"`
}

