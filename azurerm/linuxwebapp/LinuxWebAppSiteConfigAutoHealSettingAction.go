package linuxwebapp


type LinuxWebAppSiteConfigAutoHealSettingAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#action_type LinuxWebApp#action_type}.
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#minimum_process_execution_time LinuxWebApp#minimum_process_execution_time}.
	MinimumProcessExecutionTime *string `field:"optional" json:"minimumProcessExecutionTime" yaml:"minimumProcessExecutionTime"`
}

