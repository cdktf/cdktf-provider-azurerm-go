package logicappworkflow


type LogicAppWorkflowAccessControlTrigger struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/logic_app_workflow#allowed_caller_ip_address_range LogicAppWorkflow#allowed_caller_ip_address_range}.
	AllowedCallerIpAddressRange *[]*string `field:"required" json:"allowedCallerIpAddressRange" yaml:"allowedCallerIpAddressRange"`
	// open_authentication_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/logic_app_workflow#open_authentication_policy LogicAppWorkflow#open_authentication_policy}
	OpenAuthenticationPolicy interface{} `field:"optional" json:"openAuthenticationPolicy" yaml:"openAuthenticationPolicy"`
}

