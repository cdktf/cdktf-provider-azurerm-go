package logicappworkflow


type LogicAppWorkflowAccessControlAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_workflow#allowed_caller_ip_address_range LogicAppWorkflow#allowed_caller_ip_address_range}.
	AllowedCallerIpAddressRange *[]*string `field:"required" json:"allowedCallerIpAddressRange" yaml:"allowedCallerIpAddressRange"`
}

