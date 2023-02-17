package monitoractiongroup


type MonitorActionGroupAzureAppPushReceiver struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_group#email_address MonitorActionGroup#email_address}.
	EmailAddress *string `field:"required" json:"emailAddress" yaml:"emailAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_action_group#name MonitorActionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

