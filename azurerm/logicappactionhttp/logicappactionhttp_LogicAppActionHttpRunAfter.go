package logicappactionhttp


type LogicAppActionHttpRunAfter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_action_http#action_name LogicAppActionHttp#action_name}.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_action_http#action_result LogicAppActionHttp#action_result}.
	ActionResult *string `field:"required" json:"actionResult" yaml:"actionResult"`
}
